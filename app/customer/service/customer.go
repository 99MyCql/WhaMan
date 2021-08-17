package service

import (
	"WhaMan/app/customer/model"
)

type Customer interface {
	Create(*model.Params) error
	Find(id uint) (*model.Customer, error)
	List() ([]*model.Customer, error)
	Update(id uint, p *model.Params) error
	Delete(id uint) error
}
