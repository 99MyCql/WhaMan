package service

import (
	"WhaMan/app/supplier/model"
)

type Supplier interface {
	// Create 新增供应商
	Create(*model.Supplier) error
	// Find 查询指定供应商
	Find(id uint) (*model.Supplier, error)
	// List 查询所有供应商，可指定查询条件和排序规则
	List() ([]*model.Supplier, error)
	// Update 更新指定供应商
	Update(id uint, supplier *model.Supplier) error
	// Delete 删除指定供应商
	Delete(id uint) error
}
