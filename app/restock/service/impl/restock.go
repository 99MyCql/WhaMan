package impl

import (
	"WhaMan/app/restock/model"
	stockModel "WhaMan/app/stock/model"
	stockService "WhaMan/app/stock/service"
	supplierModel "WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RestockImpl struct {
	stockService stockService.Stock
}

func New(stockService stockService.Stock) *RestockImpl {
	return &RestockImpl{stockService: stockService}
}

// Restock 1.新增库存和进货订单；2.更新供应商信息。
func (r *RestockImpl) Restock(p *model.RestockParams) error {
	// 执行事务：保证新增库存和进货订单、更新供应商信息等操作同时成功
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 新增库存和进货订单
		restockOrder := p.GenRestockOrder()
		stock := &stockModel.Stock{
			ModelNum:        p.ModelNum,
			Specification:   p.Specification,
			RestockQuantity: p.Quantity,
			CurQuantity:     p.Quantity,
			UnitPrice:       p.UnitPrice,
			SumMoney:        p.Quantity * p.UnitPrice,
			Location:        p.Location,
			RestockOrder:    restockOrder,
		}
		if err := tx.Create(stock).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，新增库存和进货订单出错：%+v", stock)
		}

		// 更新关联的供应商信息
		if err := r.updateSupplier(tx, restockOrder.SupplierID, restockOrder.SumMoney); err != nil {
			return errors.WithMessagef(err, "进货流程中，更新供应商信息出错：%d", restockOrder.SupplierID)
		}
		return nil
	})
}

// Find 查找
func (RestockImpl) Find(id uint) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := global.DB.Model(&model.RestockOrder{}).
		Select("restock_orders.*, suppliers.name as supplier_name").
		Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id").
		Where("restock_orders.id = ?", id).
		Scan(&data).Error
	if err != nil {
		return nil, errors.Wrapf(err, "通过ID查询进货订单出错：%d", id)
	}
	return data, nil
}

// List 查询所有进货订单，可指定查询条件和排序规则
func (RestockImpl) List(option *model.ListOption) ([]*model.RestockOrder, error) {
	tx := global.DB
	if option.Where != nil {
		if option.Where.Date != nil {
			tx = tx.Where("date >= ? and date < ?", option.Where.Date.StartDate, option.Where.Date.EndDate)
		}
		if option.Where.SupplierID != 0 {
			tx = tx.Where("supplier_id = ?", option.Where.SupplierID)
		}
		if option.Where.StockID != 0 {
			tx = tx.Where("stock_id = ?", option.Where.StockID)
		}
	}

	if option.OrderBy != "" {
		tx = tx.Order(option.OrderBy)
	}

	var restockOrders []*model.RestockOrder
	if err := tx.Find(&restockOrders).Error; err != nil {
		return nil, errors.Wrapf(err, "查询进货订单列表出错：%+v", option)
	}
	return restockOrders, nil
}

// Update 1.更新进货订单；2.更新关联的库存；3.更新关联的出货订单；4.更新关联的供应商信息
func (r *RestockImpl) Update(id uint, p *model.UpdateParams) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 查询原进货订单
		var oldRO *model.RestockOrder
		if err := tx.Find(&oldRO, id).Error; err != nil {
			return errors.WithMessagef(err, "更新进货订单过程中，查询进货订单出错：%d", id)
		}

		// 更新进货订单
		newRO := p.GenRestockOrder()
		newRO.Model = oldRO.Model
		if err := tx.Select("*").Omit("StockID").Updates(newRO).Error; err != nil {
			return errors.Wrapf(err, "更新进货订单出错：%d-%+v", id, newRO)
		}

		// 更新关联的库存
		var stock *stockModel.Stock
		if err := tx.First(&stock, oldRO.StockID).Error; err != nil {
			return errors.Wrapf(err, "更新进货订单过程中，查询库存出错：%d", oldRO.StockID)
		}
		stock.ModelNum = newRO.ModelNum
		stock.Specification = newRO.Specification
		stock.UnitPrice = newRO.UnitPrice
		stock.RestockQuantity = newRO.Quantity
		stock.CurQuantity = stock.RestockQuantity - stock.SellQuantity
		stock.SumMoney = stock.CurQuantity * stock.UnitPrice
		if err := tx.Save(stock).Error; err != nil {
			return errors.Wrapf(err, "更新进货订单过程中，更新库存出错：%+v", stock)
		}

		// 如果单价更新，更新关联的出货订单
		if oldRO.UnitPrice != newRO.UnitPrice {
			if err := r.stockService.UpdateSellOrders(tx, oldRO.StockID, newRO.UnitPrice); err != nil {
				return errors.WithMessagef(err, "更新进货订单过程中，更新库存关联的出货订单出错：%d", oldRO.SupplierID)
			}
		}

		// 如果供应商变更，更新旧新供应商信息
		if oldRO.SupplierID != newRO.SupplierID {
			if err := r.updateSupplier(tx, oldRO.SupplierID, -oldRO.SumMoney); err != nil {
				return errors.WithMessagef(err, "更新进货订单过程中，更新原供应商出错：%d", oldRO.SupplierID)
			}
			if err := r.updateSupplier(tx, newRO.SupplierID, newRO.SumMoney); err != nil {
				return errors.WithMessagef(err, "更新进货订单过程中，更新新供应商出错：%d", newRO.SupplierID)
			}
		} else if oldRO.SumMoney != newRO.SumMoney {
			// 如果总金额变更，更新供应商交易额
			if err := r.updateSupplier(tx, newRO.SupplierID, newRO.SumMoney-oldRO.SumMoney); err != nil {
				return errors.WithMessagef(err, "更新进货订单过程中，更新供应商出错：%d", oldRO.SupplierID)
			}
		}
		return nil
	})
}

// Delete 1.删除进货订单；2.删除库存；3.更新供应商信息
func (r *RestockImpl) Delete(id uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 查询进货订单
		var restockOrder *model.RestockOrder
		if err := tx.Find(&restockOrder, id).Error; err != nil {
			return errors.WithMessagef(err, "删除进货订单过程中，查询进货订单出错：%d", id)
		}
		// 删除进货订单
		if err := tx.Delete(&restockOrder).Error; err != nil {
			return errors.Wrapf(err, "删除进货订单过程中，删除进货订单出错：%d", id)
		}

		// 已出货的库存不能删除
		var stock *stockModel.Stock
		if err := tx.First(&stock, restockOrder.StockID).Error; err != nil {
			return errors.WithMessagef(err, "删除进货订单过程中，查询库存出错：%d", id)
		}
		if stock.SellQuantity != 0 {
			return errors.WithMessagef(global.ErrCannotDelete, "删除进货订单过程中，已出货的库存不能删除：%d", id)
		}
		// 删除库存
		if err := tx.Delete(&stockModel.Stock{}, restockOrder.StockID).Error; err != nil {
			return errors.Wrapf(err, "删除进货订单过程中，删除库存出错：%d", restockOrder.StockID)
		}

		// 更新关联的供应商
		if err := r.updateSupplier(tx, restockOrder.SupplierID, -restockOrder.SumMoney); err != nil {
			return errors.WithMessagef(err, "删除进货订单过程中，更新供应商出错：%d", restockOrder.SupplierID)
		}

		return nil
	})
}

// updateSupplier 更新关联的供应商
func (r RestockImpl) updateSupplier(tx *gorm.DB, supplierID uint, money float64) error {
	var supplier *supplierModel.Supplier
	if err := tx.First(&supplier, supplierID).Error; err != nil {
		return errors.Wrapf(err, "更新关联供应商的过程中，查询供应商出错：%d", supplierID)
	}
	supplier.Turnover = supplier.Turnover + money
	if err := tx.Save(supplier).Error; err != nil {
		return errors.Wrapf(err, "更新关联供应商的过程中，更新供应商出错：%+v", supplier)
	}
	return nil
}
