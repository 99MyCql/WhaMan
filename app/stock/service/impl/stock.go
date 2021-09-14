package impl

import (
	sellModel "WhaMan/app/sell/model"
	"WhaMan/app/stock/model"
	"WhaMan/pkg/global"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type StockImpl struct{}

// Find 查找
func (StockImpl) Find(id uint) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := global.DB.Model(&model.Stock{}).
		Select("stocks.*, restock_orders.date as restock_date, suppliers.name as supplier_name").
		Joins("JOIN restock_orders ON restock_orders.stock_id = stocks.id").
		Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id ").
		Where("stocks.id = ?", id).
		Scan(&data).Error
	if err != nil {
		return nil, errors.Wrapf(err, "通过ID查询库存出错：%d", id)
	}
	global.Log.Debugf("%+v", data)
	return data, nil
}

// List 获取列表
func (StockImpl) List(option *model.ListOption) ([]map[string]interface{}, error) {
	data := make([]map[string]interface{}, 0)
	tx := global.DB.Model(&model.Stock{})

	if option.Where != nil {
		if option.Where.CurQuantity != nil {
			if option.Where.CurQuantity.Start != nil {
				tx = tx.Where("cur_quantity >= ?", *option.Where.CurQuantity.Start)
			}
			if option.Where.CurQuantity.End != nil {
				tx = tx.Where("cur_quantity <= ?", *option.Where.CurQuantity.End)
			}
		}
	}

	if option.OrderBy != "" {
		tx = tx.Order(option.OrderBy)
	}
	err := tx.Select("stocks.*, restock_orders.date as restock_date, suppliers.name as supplier_name").
		Joins("JOIN restock_orders ON restock_orders.stock_id = stocks.id").
		Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id ").
		Scan(&data).Error
	if err != nil {
		return data, errors.Wrapf(err, "获取库存列表失败")
	}
	return data, nil
}

// Update
// 仅更新备注、存放地点：1.更新
// 更新数据较多时：1.更新对应的进货订单；2.更新对应的出货订单(单价)；3.重新计算当前库存和金额；4.更新库存
func (StockImpl) Update(id uint, p *model.UpdateParams) error {
	err := global.DB.Where("id = ?", id).Select("location", "note").Updates(p.GenStock()).Error
	if err != nil {
		return errors.Wrapf(err, "更新库存信息失败：%d-%+v", id, p)
	}
	return nil
}

// UpdateSellOrders 当单价发生变化时，更新关联出货订单的利润
func (StockImpl) UpdateSellOrders(tx *gorm.DB, id uint, unitPrice float64) error {
	var sellOrders []*sellModel.SellOrder
	if err := tx.Where("stock_id = ?", id).Find(&sellOrders).Error; err != nil {
		return errors.Wrapf(err, "更新关联的出货订单过程中，查询出货订单出错：%d", id)
	}
	for i := 0; i < len(sellOrders); i++ {
		sellOrders[i].RestockUnitPrice = unitPrice
		sellOrders[i].CalProfit()
		if err := tx.Model(&sellOrders[i]).Update("profit", sellOrders[i].Profit).Error; err != nil {
			return errors.Wrapf(err, "更新关联的出货订单过程中，更新出货订单出错：%+v", sellOrders[i])
		}
	}
	return nil
}
