package sell

import (
	restockDO "WhaMan/app/restock/do"
	"WhaMan/app/sell/do"
	"WhaMan/app/sell/dto"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"gorm.io/gorm"
)

type Service struct{}

// Create 新增
func (s *Service) Create(p *dto.ComReq) (uint, error) {
	sellOrder := p.Convert2SellOrder()
	log.Logger.Infof("sellOrder: %+v", sellOrder)
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 查询进货单价
		if sellOrder.RestockOrderID != nil {
			if err := tx.Model(&restockDO.RestockOrder{}).Select("unit_price").
				Where("id = ?", sellOrder.RestockOrderID).
				First(&sellOrder.RestockUnitPrice).Error; err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}
		// 新增出货订单
		if err := tx.Create(sellOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return sellOrder.ID, err
}

// Get 查找
func (Service) Get(id uint) (*dto.ComRsp, error) {
	data := &dto.ComRsp{}
	if err := database.DB.Model(&do.SellOrder{}).
		Select("sell_orders.*, customers.name as customer_name").
		Joins("JOIN customers ON sell_orders.customer_id = customers.id").
		Where("sell_orders.id = ?", id).First(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// List 查询所有出货订单，可指定查询条件和排序规则
func (Service) List(req *dto.ListReq) ([]*dto.ComRsp, error) {
	tx := database.DB.Model(&do.SellOrder{}).
		Select("sell_orders.*, customers.name as customer_name").
		Joins("JOIN customers ON sell_orders.customer_id = customers.id")
	if req.Where != nil {
		if req.Where.Date != nil {
			tx = tx.Where("date >= ? and date < ?", req.Where.Date.StartDate, req.Where.Date.EndDate)
		}
		if req.Where.CustomerID != 0 {
			tx = tx.Where("customer_id = ?", req.Where.CustomerID)
		}
		if req.Where.RestockOrderID != 0 {
			tx = tx.Where("restock_order_id = ?", req.Where.RestockOrderID)
		}
	}
	if req.OrderBy != "" {
		tx = tx.Order(req.OrderBy)
	}

	var data []*dto.ComRsp
	if err := tx.Scan(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// Update 更新
func (s *Service) Update(id uint, p *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 先查询原出货订单
		var oldSO *do.SellOrder
		if err := tx.First(&oldSO, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		log.Logger.Infof("oldSO: %+v", oldSO)

		// 更新出货订单
		newSO := p.Convert2SellOrder()
		newSO.ID = id
		// 查询进货单价
		if newSO.RestockOrderID != nil {
			if err := tx.Model(&restockDO.RestockOrder{}).Select("unit_price").
				Where("id = ?", newSO.RestockOrderID).
				First(&newSO.RestockUnitPrice).Error; err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}
		// 更新
		log.Logger.Infof("newSO: %+v", newSO)
		if err := tx.Select("*").Omit("CreatedAt").Updates(newSO).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// Delete 1.更新库存；2.更新客户；3.删除出货订单
func (s *Service) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 删除出货订单（软删除）
		if err := tx.Delete(&do.SellOrder{}, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}
