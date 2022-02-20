package restock

import (
	"WhaMan/app/restock/do"
	"WhaMan/app/restock/dto"
	_stock "WhaMan/app/stock"
	stockDO "WhaMan/app/stock/do"
	supplierDO "WhaMan/app/supplier/do"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var stockService = new(_stock.Service)

type Service struct{}

// Create 创建进货
func (s *Service) Create(p *dto.ComReq) (uint, error) {
	stock := p.Convert2Stock()
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 新增库存和进货订单
		log.Logger.Infof("stock: %+v", stock)
		if err := tx.Create(&stock).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		restockOrder := p.Convert2RestockOrder()
		restockOrder.StockID = stock.ID
		log.Logger.Infof("restockOrder: %+v", restockOrder)
		if err := tx.Create(&restockOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 更新关联的供应商信息
		if err := s.updateSupplier(tx, restockOrder.SupplierID, restockOrder.SumMoney,
			restockOrder.PaidMoney); err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return stock.ID, err
}

// Get 查找
func (Service) Get(id uint) (*dto.ComRsp, error) {
	data := &dto.ComRsp{}
	err := database.DB.Model(&do.RestockOrder{}).
		Select("restock_orders.*, suppliers.name as supplier_name").
		Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id").
		Where("restock_orders.id = ?", id).
		Scan(&data).Error
	if err != nil {
		return nil, myErr.ServerErr
	}
	return data, nil
}

// List 查询所有进货订单，可指定查询条件和排序规则
func (Service) List(req *dto.ListReq) ([]*dto.ComRsp, error) {
	tx := database.DB.Model(&do.RestockOrder{})
	if req.Where != nil {
		if req.Where.Date != nil {
			tx = tx.Where("date >= ? and date < ?", req.Where.Date.StartDate, req.Where.Date.EndDate)
		}
		if req.Where.SupplierID != 0 {
			tx = tx.Where("supplier_id = ?", req.Where.SupplierID)
		}
		if req.Where.StockID != 0 {
			tx = tx.Where("stock_id = ?", req.Where.StockID)
		}
	}
	if req.OrderBy != "" {
		tx = tx.Order(req.OrderBy)
	}
	data := make([]*dto.ComRsp, 0)
	if err := tx.Scan(&data).Error; err != nil {
		return nil, myErr.ServerErr
	}
	return data, nil
}

// Update 1.更新进货订单；2.更新关联的库存；3.更新关联的出货订单；4.更新关联的供应商信息
func (s *Service) Update(id uint, p *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 查询原进货订单
		var oldRO *do.RestockOrder
		if err := tx.Find(&oldRO, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		log.Logger.Infof("oldRO: %+v", oldRO)

		// 更新进货订单
		newRO := p.Convert2RestockOrder()
		newRO.ID = id
		log.Logger.Infof("newRO: %+v", newRO)
		if err := tx.Select("*").Omit("StockID", "CreatedAt").Updates(&newRO).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}

		// 更新关联的库存
		var stock *stockDO.Stock
		if err := tx.Where("id = ?", newRO.StockID).First(&stock).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		stock.ModelNum = newRO.ModelNum
		stock.Specification = newRO.Specification
		stock.UnitPrice = newRO.UnitPrice
		stock.RestockQuantity = newRO.Quantity
		stock.CurQuantity = stock.RestockQuantity - stock.SellQuantity
		stock.SumMoney = stock.CurQuantity * stock.UnitPrice
		stock.Location = newRO.Location
		stock.Note = newRO.Note
		log.Logger.Infof("stock: %+v", stock)
		if err := tx.Save(stock).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}

		// 如果单价更新，更新关联的出货订单
		if oldRO.UnitPrice != newRO.UnitPrice {
			if err := stockService.UpdateSellOrders(tx, oldRO.StockID, newRO.UnitPrice); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}

		if oldRO.SupplierID != newRO.SupplierID {
			// 如果供应商变更，更新旧、新供应商信息
			if err := s.updateSupplier(tx, oldRO.SupplierID, -oldRO.SumMoney, -oldRO.PaidMoney); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
			if err := s.updateSupplier(tx, newRO.SupplierID, newRO.SumMoney, newRO.PaidMoney); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		} else if oldRO.SumMoney != newRO.SumMoney || oldRO.PaidMoney != newRO.PaidMoney {
			// 如果供应商未变更，但总金额或已付金额变更，更新供应商信息
			if err := s.updateSupplier(tx, newRO.SupplierID, newRO.SumMoney-oldRO.SumMoney,
				newRO.PaidMoney-oldRO.PaidMoney); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}
		return nil
	})
}

// Delete 1.删除进货订单；2.删除库存；3.更新供应商信息
func (s *Service) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 查询进货订单
		var restockOrder *do.RestockOrder
		if err := tx.Where("id = ?", id).Find(&restockOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 查询关联的库存
		var stock *stockDO.Stock
		if err := tx.Where("id = ?", restockOrder.StockID).First(&stock).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 已出货的库存不能删除
		if stock.SellQuantity != 0 {
			return myErr.CannotDelete.AddMsg("已出货的库存不能删除")
		}
		// 删除库存
		if err := tx.Delete(&stockDO.Stock{}, restockOrder.StockID).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 删除进货订单
		if err := tx.Delete(&restockOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}

		// 更新关联的供应商
		if err := s.updateSupplier(tx, restockOrder.SupplierID, -restockOrder.SumMoney,
			-restockOrder.PaidMoney); err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// updateSupplier 更新关联的供应商
func (Service) updateSupplier(tx *gorm.DB, supplierID uint, sumMoney float64, paidMoney float64) error {
	var supplier *supplierDO.Supplier
	if err := tx.Where("id = ?", supplierID).First(&supplier).Error; err != nil {
		return errors.Wrapf(err, "更新关联供应商的过程中，查询供应商出错：%d", supplierID)
	}
	supplier.Turnover = supplier.Turnover + sumMoney
	supplier.UnpaidMoney += sumMoney - paidMoney
	if err := tx.Save(supplier).Error; err != nil {
		return errors.Wrapf(err, "更新关联供应商的过程中，更新供应商出错：%+v", supplier)
	}
	return nil
}
