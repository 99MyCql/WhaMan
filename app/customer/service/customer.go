package service

import (
	"WhaMan/app/customer/model"
	sellModel "WhaMan/app/sell/model"
)

type Customer interface {
	Create(i *model.CustomerInfo) error
	Find(name string) (*model.Customer, error)
	List() ([]*model.Customer, error)
	ListSellOrders(name string) ([]*sellModel.SellOrder, error)
	Update(i *model.CustomerInfo) error
	Delete(name string) error
}
