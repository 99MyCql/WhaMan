package impl

import (
	customerModel "WhaMan/app/customer/model"
	"WhaMan/app/sell/model"
	stockModel "WhaMan/app/stock/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type SellImpl struct{}

// Sell 1.新增出货订单；2.更新库存；3.更新客户
func (s *SellImpl) Sell(p *model.Params) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 新增出货订单
		sellOrder := p.GenSellOrder()
		sellOrder.CalProfit()
		if err := tx.Create(sellOrder).Error; err != nil {
			return errors.Wrapf(err, "出货流程中，新增出货订单出错：%+v", sellOrder)
		}

		// 更新库存
		if err := s.updateStock(tx, p.StockID, p.Quantity); err != nil {
			return errors.WithMessagef(err, "出货流程中，更新关联库存出错：%+v", p)
		}

		// 更新客户
		if err := s.updateCustomer(tx, sellOrder.CustomerID, sellOrder.SumMoney, sellOrder.PaidMoney); err != nil {
			return errors.WithMessagef(err, "出货流程中，更新客户信息出错：%d", sellOrder.CustomerID)
		}
		return nil
	})
}

// Find 查找
func (SellImpl) Find(id uint) (*model.SellOrder, error) {
	// TODO: 连接查询获取客户信息和库存信息
	var sellOrder *model.SellOrder
	if err := global.DB.Find(&sellOrder, id).Error; err != nil {
		return nil, errors.Wrapf(err, "通过ID查询出货订单出错：%d", id)
	}
	return sellOrder, nil
}

// List 查询所有出货订单，可指定查询条件和排序规则
func (SellImpl) List(option *model.ListOption) ([]*model.SellOrder, error) {
	tx := global.DB
	if option.Where != nil {
		if option.Where.Date != nil {
			tx = tx.Where("date >= ? and date < ?", option.Where.Date.StartDate, option.Where.Date.EndDate)
		}
		if option.Where.CustomerID != 0 {
			tx = tx.Where("customer_id = ?", option.Where.CustomerID)
		}
	}

	if option.OrderBy != "" {
		tx = tx.Order(option.OrderBy)
	}

	var sellOrders []*model.SellOrder
	if err := tx.Find(&sellOrders).Error; err != nil {
		return nil, errors.Wrapf(err, "查询出货订单列表出错：%+v", option)
	}
	return sellOrders, nil
}

// Update 1.更新出货订单；2.更新库存；3.更新客户
func (s *SellImpl) Update(id uint, p *model.Params) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var oldSO *model.SellOrder
		if err := tx.First(&oldSO, id).Error; err != nil {
			return errors.Wrapf(err, "更新出货订单过程中，查询原库存出错：%d", id)
		}

		// 更新出货订单
		newSO := p.GenSellOrder()
		newSO.Model = oldSO.Model
		newSO.CalProfit()
		if err := tx.Save(newSO).Error; err != nil {
			return errors.Wrapf(err, "更新出货订单出错：%d-%+v", id, newSO)
		}

		// 更新库存
		if oldSO.StockID != newSO.StockID {
			// 若库存变更，同时更新新旧库存
			if err := s.updateStock(tx, oldSO.StockID, -oldSO.Quantity); err != nil {
				return errors.WithMessagef(err, "更新出货订单过程中，更新原库存出错：%+v", oldSO)
			}
			if err := s.updateStock(tx, newSO.StockID, newSO.Quantity); err != nil {
				return errors.WithMessagef(err, "更新出货订单过程中，更新新库存出错：%+v", newSO)
			}
		} else if oldSO.Quantity != newSO.Quantity {
			if err := s.updateStock(tx, newSO.StockID, newSO.Quantity-oldSO.Quantity); err != nil {
				return errors.WithMessagef(err, "更新出货订单过程中，更新库存出错：%+v", newSO)
			}
		}

		// 更新客户
		if oldSO.CustomerID != newSO.CustomerID {
			if err := s.updateCustomer(tx, oldSO.CustomerID, -oldSO.SumMoney, -oldSO.PaidMoney); err != nil {
				return errors.WithMessagef(err, "更新出货订单过程中，更新原客户信息出错：%d", oldSO.CustomerID)
			}
			if err := s.updateCustomer(tx, newSO.CustomerID, newSO.SumMoney, newSO.PaidMoney); err != nil {
				return errors.WithMessagef(err, "更新出货订单过程中，更新新客户信息出错：%d", newSO.CustomerID)
			}
		} else if oldSO.SumMoney != newSO.SumMoney || oldSO.PaidMoney != newSO.PaidMoney {
			if err := s.updateCustomer(tx, newSO.CustomerID,
				newSO.SumMoney-oldSO.SumMoney, newSO.PaidMoney-oldSO.PaidMoney); err != nil {
				return errors.WithMessagef(err, "更新出货订单过程中，更新客户信息出错：%d", newSO.CustomerID)
			}
		}

		return nil
	})
}

// Delete 1.更新库存；2.更新客户；3.删除出货订单
func (s *SellImpl) Delete(id uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var sellOrder *model.SellOrder
		if err := tx.Find(&sellOrder, id).Error; err != nil {
			return errors.Wrapf(err, "删除出货订单过程中，查询出货订单出错：%d", id)
		}

		// 更新库存
		if err := s.updateStock(tx, sellOrder.StockID, -sellOrder.Quantity); err != nil {
			return errors.WithMessagef(err, "删除出货订单过程中，更新库存出错：%+v", sellOrder)
		}

		// 更新客户
		if err := s.updateCustomer(tx, sellOrder.CustomerID, -sellOrder.SumMoney, -sellOrder.PaidMoney); err != nil {
			return errors.WithMessagef(err, "删除出货订单过程中，更新客户信息出错：%+v", sellOrder)
		}

		// 删除出货订单
		if err := tx.Delete(&sellOrder).Error; err != nil {
			return errors.Wrapf(err, "删除出货订单过程中，删除出货订单出错：%d", id)
		}

		return nil
	})
}

// updateCustomer 更新关联的客户
func (SellImpl) updateCustomer(tx *gorm.DB, customerID uint, sumMoney float64, paidMoney float64) error {
	var customer *customerModel.Customer
	if err := tx.First(&customer, customerID).Error; err != nil {
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
func (s SellImpl) updateStock(tx *gorm.DB, stockID *uint, quantity float64) error {
	if stockID == nil {
		return nil
	}

	var stock *stockModel.Stock
	if err := tx.First(&stock, *stockID).Error; err != nil {
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
