package impl

import (
	"WhaMan/app/stock/model"
	"WhaMan/pkg/global"

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
// 仅更新备注、存放地点：1.检查更新权限；2.更新
// 更新数据较多时：1.检查更新权限；2.更新对应的进货订单；3.更新对应的出货订单(单价)；4.重新计算当前库存和金额；5.更新库存
func (StockImpl) Update(id uint, p *model.UpdateParams) error {
	// TODO: 检查更新权限

	if err := global.DB.Where("id = ?", id).Updates(p.GenStock()).Error; err != nil {
		return errors.Wrapf(err, "更新库存信息失败：%d-%+v", id, p)
	}
	return nil
}

// Delete 1.检查删除权限；2.已出货的库存不能删除；3.删除库存(由于设置了OnDelete:CASCADE，会同时删除进货订单)
func (s *StockImpl) Delete(id uint) error {
	// TODO: 检查删除权限

	// 已出货的库存不能删除
	stock, err := s.Find(id)
	if err != nil {
		return errors.WithMessagef(err, "删除客户过程中，根据ID查询出错：%d", id)
	}
	if stock.SellQuantity != 0 {
		return errors.WithMessagef(global.ErrCannotDelete, "删除库存过程中，已出货的库存不能删除：%d", id)
	}

	if err := global.DB.Delete(&model.Stock{}, id).Error; err != nil {
		return errors.Wrapf(err, "删除库存出错：%d", id)
	}
	return nil
}
