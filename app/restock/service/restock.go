package service

import (
	"WhaMan/app/restock/model"
)

// Restock 进货模块
type Restock interface {
	Restock(*model.RestockParams) error
	Find(id uint) (map[string]interface{}, error)
	List(option *model.ListOption) ([]*model.RestockOrder, error)
	Update(id uint, p *model.RestockParams) error
	Delete(id uint) error
}
