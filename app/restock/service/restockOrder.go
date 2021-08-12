package service

import (
	"WhaMan/app/restock/model"
)

type RestockOrder interface {
	Create(r *model.RestockOrder) error
	Find(id string) (*model.RestockOrder, error)
	List() ([]*model.RestockOrder, error)
	Update(r *model.RestockOrder) error
	Delete(id string) error
}
