package service

import (
	"WhaMan/app/restock/model"
)

// Restock 进货模块
type Restock interface {
	// Restock 进货：新增进货订单，并新增库存，以及更新供应商交易额信息
	Restock(i *model.RestockInfo) error
	// Find 查询指定进货订单
	Find(id string) (*model.RestockOrder, error)
	// List 查询所有进货订单，可指定查询条件和排序规则
	List() ([]*model.RestockOrder, error)
	// Update 更新指定进货订单，并更新关联的库存和供应商信息
	Update(i *model.RestockInfo) error
	// Delete 删除指定进货订单，删除关联的库存，更新关联的供应商信息
	Delete(id string) error
}
