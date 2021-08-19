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
func (StockImpl) Find(id uint) (*model.Stock, error) {
	var stock *model.Stock
	if err := global.DB.First(&stock, id).Error; err != nil {
		return nil, errors.Wrapf(err, "通过ID查询库存出错：%d", id)
	}
	return stock, nil
}

// List 获取列表
func (StockImpl) List() ([]*model.Stock, error) {
	var stocks []*model.Stock
	if err := global.DB.Find(&stocks).Error; err != nil {
		return nil, errors.Wrapf(err, "获取库存列表失败")
	}
	return stocks, nil
}

// Update
// 仅更新备注、存放地点：1.更新
// 更新数据较多时：1.更新对应的进货订单；2.更新对应的出货订单(单价)；3.重新计算当前库存和金额；4.更新库存
func (StockImpl) Update(id uint, p *model.UpdateParams) error {
	if err := global.DB.Where("id = ?", id).Updates(p.GenStock()).Error; err != nil {
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
		sellOrders[i].CalProfit(unitPrice)
		if err := tx.Model(&sellOrders[i]).Update("profit", sellOrders[i].Profit).Error; err != nil {
			return errors.Wrapf(err, "更新关联的出货订单过程中，更新出货订单出错：%+v", sellOrders[i])
		}
	}
	return nil
}
