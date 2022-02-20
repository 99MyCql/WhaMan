package sell

import (
	customerDO "WhaMan/app/customer/do"
	"WhaMan/app/sell/do"
	"WhaMan/app/sell/dto"
	stockDO "WhaMan/app/stock/do"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Service struct{}

// Create 新增
func (s *Service) Create(p *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 新增出货订单
		sellOrder := p.Convert2SellOrder()
		log.Logger.Infof("sellOrder: %+v", sellOrder)
		if err := tx.Create(sellOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 更新关联库存
		if err := s.updateStock(tx, p.StockID, p.Quantity); err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 更新关联客户
		if err := s.updateCustomer(tx, sellOrder.CustomerID, sellOrder.SumMoney, sellOrder.PaidMoney); err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// Get 查找
func (Service) Get(id uint) (*dto.ComRsp, error) {
	data := &dto.ComRsp{}
	err := database.DB.Model(&do.SellOrder{}).
		Select("sell_orders.*, customers.name as customer_name").
		Joins("JOIN customers ON sell_orders.customer_id = customers.id").
		Where("sell_orders.id = ?", id).
		Scan(&data).Error
	if err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// List 查询所有出货订单，可指定查询条件和排序规则
func (Service) List(req *dto.ListReq) ([]*dto.ComRsp, error) {
	tx := database.DB.Model(&do.SellOrder{})
	if req.Where != nil {
		if req.Where.Date != nil {
			tx = tx.Where("date >= ? and date < ?", req.Where.Date.StartDate, req.Where.Date.EndDate)
		}
		if req.Where.CustomerID != 0 {
			tx = tx.Where("customer_id = ?", req.Where.CustomerID)
		}
		if req.Where.StockID != 0 {
			tx = tx.Where("stock_id = ?", req.Where.StockID)
		}
	}
	if req.OrderBy != "" {
		tx = tx.Order(req.OrderBy)
	}

	var data []*dto.ComRsp
	if err := tx.Find(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// Update 1.更新出货订单；2.更新库存；3.更新客户
func (s *Service) Update(id uint, p *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var oldSO *do.SellOrder
		if err := tx.First(&oldSO, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		log.Logger.Infof("oldSO: %+v", oldSO)

		// 更新出货订单
		newSO := p.Convert2SellOrder()
		newSO.ID = id
		log.Logger.Infof("newSO: %+v", newSO)
		// stock_id 为 0 时，Save更新数据会出错
		var err error
		if newSO.StockID == 0 {
			err = tx.Select("*").Omit("StockID", "CreatedAt").Updates(newSO).Error
		} else {
			err = tx.Save(newSO).Error
		}
		if err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}

		// 更新库存
		if oldSO.StockID != newSO.StockID {
			// 若库存变更，同时更新新旧库存
			if err := s.updateStock(tx, oldSO.StockID, -oldSO.Quantity); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
			if err := s.updateStock(tx, newSO.StockID, newSO.Quantity); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		} else if oldSO.Quantity != newSO.Quantity {
			if err := s.updateStock(tx, newSO.StockID, newSO.Quantity-oldSO.Quantity); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}

		// 更新客户
		if oldSO.CustomerID != newSO.CustomerID {
			if err := s.updateCustomer(tx, oldSO.CustomerID, -oldSO.SumMoney, -oldSO.PaidMoney); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
			if err := s.updateCustomer(tx, newSO.CustomerID, newSO.SumMoney, newSO.PaidMoney); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		} else if oldSO.SumMoney != newSO.SumMoney || oldSO.PaidMoney != newSO.PaidMoney {
			if err := s.updateCustomer(tx, newSO.CustomerID,
				newSO.SumMoney-oldSO.SumMoney, newSO.PaidMoney-oldSO.PaidMoney); err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
		}
		return nil
	})
}

// Delete 1.更新库存；2.更新客户；3.删除出货订单
func (s *Service) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var sellOrder *do.SellOrder
		if err := tx.Find(&sellOrder, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		log.Logger.Infof("%+v", sellOrder)
		// 更新库存
		if err := s.updateStock(tx, sellOrder.StockID, -sellOrder.Quantity); err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 更新客户
		if err := s.updateCustomer(tx, sellOrder.CustomerID, -sellOrder.SumMoney, -sellOrder.PaidMoney); err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		// 删除出货订单
		if err := tx.Delete(&sellOrder).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// updateCustomer 更新关联的客户
func (Service) updateCustomer(tx *gorm.DB, customerID uint, sumMoney float64, paidMoney float64) error {
	var customer *customerDO.Customer
	if err := tx.Where("id = ?", customerID).First(&customer).Error; err != nil {
		return errors.Wrapf(err, "更新关联客户的过程中，查询客户出错：%d", customerID)
	}
	customer.Turnover += sumMoney
	customer.UnpaidMoney += sumMoney - paidMoney
	if err := tx.Save(customer).Error; err != nil {
		return errors.Wrapf(err, "更新关联客户的过程中，更新客户出错：%+v", customer)
	}
	return nil
}

// updateStock 更新关联的库存
func (s Service) updateStock(tx *gorm.DB, stockID uint, quantity float64) error {
	if stockID == 0 {
		return nil
	}

	var stock *stockDO.Stock
	if err := tx.Where("id = ?", stockID).First(&stock).Error; err != nil {
		return errors.Wrapf(err, "更新关联库存过程中，查询库存出错：%d", stockID)
	}
	stock.SellQuantity += quantity
	stock.CurQuantity = stock.RestockQuantity - stock.SellQuantity
	stock.SumMoney = stock.CurQuantity * stock.UnitPrice
	if err := tx.Save(stock).Error; err != nil {
		return errors.Wrapf(err, "更新关联库存过程中，更新库存出错：%+v", stock)
	}
	return nil
}
