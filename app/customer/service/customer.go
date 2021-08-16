package service

import (
	"WhaMan/app/customer/model"
)

type Customer interface {
	Create(i *model.Customer) error
	Find(id uint) (*model.Customer, error)
	List() ([]*model.Customer, error)
	Update(id uint, i *model.Customer) error
	Delete(id uint) error
	// FindByName(name string) (*model.Customer, error)
}
