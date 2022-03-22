package restock

import (
	"WhaMan/app/restock/do"
	"WhaMan/app/restock/dto"
	sellDO "WhaMan/app/sell/do"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Service struct{}

// Create 创建进货
func (s *Service) Create(p *dto.ComReq) (uint, error) {
	restockOrder := p.Convert2RestockOrder()
	log.Logger.Infof("restockOrder: %+v", restockOrder)
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&restockOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return restockOrder.ID, err
}

// Get 查找
func (Service) Get(id uint) (*dto.GetRsp, error) {
	var restockOrder *dto.GetRsp
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&do.RestockOrder{}).
			Select("restock_orders.*, suppliers.name as supplier_name").
			Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id").
			Where("restock_orders.id = ?", id).
			First(&restockOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if err := tx.Model(&sellDO.SellOrder{}).
			Select("sell_orders.*, customers.name as customer_name").
			Joins("JOIN customers ON sell_orders.customer_id = customers.id").
			Where("sell_orders.restock_order_id = ?", id).
			Order("date desc").
			Scan(&restockOrder.SellOrders).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return restockOrder, err
}

// List 查询所有进货订单，可指定查询条件和排序规则
func (Service) List(req *dto.ListReq) ([]*dto.ListRsp, error) {
	var restockOrders []*dto.ListRsp
	err := database.DB.Transaction(func(txDB *gorm.DB) error {
		subTx := txDB.Model(&do.RestockOrder{}).
			Select("restock_orders.*, suppliers.name as supplier_name, "+
				"SUM(IFNULL(so.quantity, 0)) as sell_quantity, "+
				"SUM((so.unit_price-so.restock_unit_price)*so.quantity-so.freight_cost-so.kickback-so.tax-so.other_cost) as profit").
			Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id").
			Joins("LEFT JOIN (?) so ON restock_orders.id = so.restock_order_id",
				database.DB.Model(&sellDO.SellOrder{}).
					Select("quantity, unit_price, restock_unit_price, freight_cost, "+
						"kickback, tax, other_cost, restock_order_id"),
			).
			Group("restock_orders.id")
		tx := txDB.Table("(?) t", subTx).Select("*") // 如此才能使用别名

		// 设置排序规则
		if req.OrderBy == "" {
			req.OrderBy = "date desc"
		}
		tx = tx.Order(req.OrderBy)

		// 设置查询条件
		if req.Where != nil {
			if req.Where.Date != nil {
				if req.Where.Date.StartDate != "" {
					tx = tx.Where("date >= ?", req.Where.Date.StartDate)
				}
				if req.Where.Date.EndDate != "" {
					tx = tx.Where("date < ?", req.Where.Date.EndDate)
				}
			}
			if req.Where.SupplierID != 0 {
				tx = tx.Where("supplier_id = ?", req.Where.SupplierID)
			}
			if req.Where.ModelNum != "" {
				tx = tx.Where("model_num = ?", req.Where.ModelNum)
			}
			if req.Where.CurQuantity != nil {
				if req.Where.CurQuantity.Start != nil {
					tx = tx.Where("quantity - sell_quantity >= ?", req.Where.CurQuantity.Start)
				}
				if req.Where.CurQuantity.End != nil {
					tx = tx.Where("quantity - sell_quantity < ?", req.Where.CurQuantity.End)
				}
			}
		}

		// 查询进货订单
		if err := tx.Scan(&restockOrders).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return restockOrders, err
}

// ListGroupByModelNum 通过型号进行分类，统计每个型号的剩余库存、价值金额、利润
func (Service) ListGroupByModelNum(req *dto.ListGroupByModelNumReq) ([]*dto.ListGroupByModelNumRsp, error) {
	// 构造子查询，连接查询关联的出货订单，统计每笔进货订单的剩余库存、利润
	subTx := database.DB.Model(&do.RestockOrder{}).
		Select("restock_orders.model_num, restock_orders.unit_price, "+
			"restock_orders.quantity-SUM(IFNULL(so.quantity, 0)) as cur_quantity, "+
			"SUM((so.unit_price-so.restock_unit_price)*so.quantity-so.freight_cost-so.kickback-so.tax-so.other_cost) as profit").
		Joins("LEFT JOIN (?) so ON restock_orders.id = so.restock_order_id",
			database.DB.Model(&sellDO.SellOrder{}).
				Select("quantity, unit_price, restock_unit_price, freight_cost, "+
					"kickback, tax, other_cost, restock_order_id"),
		).
		Group("restock_orders.id")

	tx := database.DB.Table("(?) as t", subTx).
		Select("t.model_num, SUM(t.cur_quantity) as cur_quantity, " +
			"SUM(t.cur_quantity*t.unit_price) as sum_money, SUM(t.profit) as profit").
		Group("t.model_num")
	if req.OrderBy == "" {
		req.OrderBy = "CONVERT(model_num USING gbk)"
	}
	tx = tx.Order(req.OrderBy)
	var data []*dto.ListGroupByModelNumRsp
	if err := tx.Scan(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// Update 更新
func (s *Service) Update(id uint, req *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 先查询原进货订单
		var oldRO *do.RestockOrder
		if err := tx.First(&oldRO, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		log.Logger.Infof("oldRO: %+v", oldRO)

		// 再更新进货订单
		newRO := req.Convert2RestockOrder()
		newRO.ID = id
		log.Logger.Infof("newRO: %+v", newRO)
		if err := tx.Select("*").Omit("CreatedAt").Updates(&newRO).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}

		// 如果单价更新，更新关联的出货订单
		if oldRO.UnitPrice != newRO.UnitPrice {
			if err := s.updateSellOrders(tx, oldRO.ID, newRO.UnitPrice); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}
		return nil
	})
}

// Delete 删除
func (s *Service) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 存在关联的出货订单，不能删除
		var sellOrders []sellDO.SellOrder
		if err := tx.Model(&sellDO.SellOrder{}).Where("restock_order_id = ?", id).
			Find(&sellOrders).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if len(sellOrders) != 0 {
			return myErr.CannotDelete.SetDetail("存在关联的出货订单，不能删除")
		}
		// 删除进货订单（软删除）
		if err := tx.Delete(&do.RestockOrder{}, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// updateSellOrders 更新关联出货订单
func (Service) updateSellOrders(tx *gorm.DB, restockOrderID uint, unitPrice float64) error {
	var sellOrders []*sellDO.SellOrder
	if err := tx.Model(&sellDO.SellOrder{}).Where("restock_order_id = ?", restockOrderID).
		Find(&sellOrders).Error; err != nil {
		return errors.Wrapf(err, "更新关联出货订单的过程中，查询出错，restockOrderID:%d", restockOrderID)
	}
	for i := 0; i < len(sellOrders); i++ {
		sellOrders[i].RestockUnitPrice = unitPrice
		if err := tx.Save(&sellOrders[i]).Error; err != nil {
			return errors.Wrapf(err, "更新关联出货订单的过程中，更新出货订单出错，sellOrder:%+v", sellOrders[i])
		}
	}
	return nil
}
