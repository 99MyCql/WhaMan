package service

import (
	"WhaMan/app/customer/model"
	sellModel "WhaMan/app/sell/model"
)

type Customer interface {
	Create(customer *model.Customer) error
	Find(name string) (*model.Customer, error)
	List() ([]*model.Customer, error)
	ListSellOrders(name string) ([]*sellModel.SellOrder, error)
	Update(customer *model.Customer) error
	Delete(name string) error
}
