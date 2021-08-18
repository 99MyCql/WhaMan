package impl

import (
	"WhaMan/app/restock/model"
	sellModel "WhaMan/app/sell/model"
	stockModel "WhaMan/app/stock/model"
	supplierModel "WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RestockImpl struct {
}

// Restock 1.新增库存和进货订单；2.更新供应商信息。
func (r *RestockImpl) Restock(p *model.RestockParams) error {
	// 执行事务：保证新增库存和进货订单、更新供应商信息等操作同时成功
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 新增库存和进货订单
		stock := &stockModel.Stock{
			ModelNum:        p.ModelNum,
			Specification:   p.Specification,
			RestockQuantity: p.Quantity,
			CurQuantity:     p.Quantity,
			UnitPrice:       p.UnitPrice,
			SumMoney:        p.SumMoney,
			Location:        p.Location,
			RestockOrder:    p.GenRestockOrder(),
		}
		if err := tx.Create(stock).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，新增库存和进货订单失败：%+v", stock)
		}

		// 更新关联的供应商信息
		var supplier supplierModel.Supplier
		if err := tx.First(&supplier, "id = ?", p.SupplierID).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，获取供应商信息失败：%d", p.SupplierID)
		}
		supplier.Turnover += p.SumMoney
		if err := tx.Save(supplier).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，更新供应商信息失败：%+v", supplier)
		}
		return nil
	})
}

// Find 查找
func (RestockImpl) Find(id uint) (*model.RestockOrder, error) {
	var restockOrder *model.RestockOrder
	if err := global.DB.Find(&restockOrder, id).Error; err != nil {
		return nil, errors.Wrapf(err, "通过ID查询进货订单出错：%d", id)
	}
	return restockOrder, nil
}

// List 查询所有进货订单，可指定查询条件和排序规则
func (RestockImpl) List(option *model.ListOption) ([]*model.RestockOrder, error) {
	tx := global.DB.Model(&model.RestockOrder{})
	if option.Where != nil {
		if option.Where.Date != nil {
			tx = tx.Where("date >= ? and date < ?", option.Where.Date.StartDate, option.Where.Date.EndDate)
		}
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
		var old *model.RestockOrder
		if err := tx.Find(&old, id).Error; err != nil {
			return errors.WithMessagef(err, "更新进货订单过程中，查询进货订单出错：%d", id)
		}

		// 更新进货订单
		if err := tx.Where("id = ?", id).Updates(p.GenRestockOrder()).Error; err != nil {
			return errors.Wrapf(err, "更新进货订单出错：%d-%+v", id, p)
		}

		// 更新关联的库存
		var stock *stockModel.Stock
		if err := tx.First(&stock, old.StockID).Error; err != nil {
			return errors.Wrapf(err, "更新进货订单过程中，查询库存出错：%d", old.StockID)
		}
		stock.ModelNum = p.ModelNum
		stock.Specification = p.Specification
		stock.UnitPrice = p.UnitPrice
		stock.RestockQuantity = p.Quantity
		stock.CurQuantity = stock.RestockQuantity - stock.SellQuantity
		stock.SumMoney = stock.CurQuantity * stock.UnitPrice
		if err := tx.Save(stock).Error; err != nil {
			return errors.Wrapf(err, "更新进货订单过程中，更新库存出错：%+v", stock)
		}

		// 如果单价更新，更新关联的出货订单
		if old.UnitPrice != p.UnitPrice {
			var sellOrders []*sellModel.SellOrder
			if err := tx.Where("stock_id = ?", old.StockID).Find(&sellOrders).Error; err != nil {
				return errors.Wrapf(err, "更新进货订单过程中，查询出货订单出错：%d", old.StockID)
			}
			for i := 0; i < len(sellOrders); i++ {
				sellOrders[i].Profit = sellOrders[i].Quantity*(sellOrders[i].UnitPrice-stock.UnitPrice) -
					sellOrders[i].FreightCost - sellOrders[i].Kickback - sellOrders[i].Tax - sellOrders[i].OtherCost
				if err := tx.Save(sellOrders[i]).Error; err != nil {
					return errors.Wrapf(err, "更新进货订单过程中，更新出货订单出错：%+v", stock)
				}
			}
		}

		// 如果供应商变更，更新供应商信息
		if old.SupplierID != p.SupplierID {
			var oldSupplier *supplierModel.Supplier
			if err := tx.First(&oldSupplier, old.SupplierID).Error; err != nil {
				return errors.Wrapf(err, "更新进货订单过程中，查询原供应商出错：%d", old.SupplierID)
			}
			oldSupplier.Turnover = oldSupplier.Turnover - old.SumMoney
			if err := tx.Save(oldSupplier).Error; err != nil {
				return errors.Wrapf(err, "更新进货订单过程中，更新原供应商出错：%+v", oldSupplier)
			}

			var supplier *supplierModel.Supplier
			if err := tx.First(&supplier, p.SupplierID).Error; err != nil {
				return errors.Wrapf(err, "更新进货订单过程中，查询新供应商出错：%d", old.SupplierID)
			}
			supplier.Turnover = supplier.Turnover + p.SumMoney
			if err := tx.Save(supplier).Error; err != nil {
				return errors.Wrapf(err, "更新进货订单过程中，更新新供应商出错：%+v", supplier)
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
		var supplier *supplierModel.Supplier
		if err := tx.First(&supplier, restockOrder.SupplierID).Error; err != nil {
			return errors.Wrapf(err, "删除进货订单过程中，查询供应商出错：%d", restockOrder.SupplierID)
		}
		supplier.Turnover = supplier.Turnover - restockOrder.SumMoney
		if err := tx.Save(supplier).Error; err != nil {
			return errors.Wrapf(err, "删除进货订单过程中，更新供应商出错：%+v", supplier)
		}

		return nil
	})
}
