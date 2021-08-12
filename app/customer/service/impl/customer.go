package impl

import (
	"WhaMan/app/customer/model"
	sellModel "WhaMan/app/sell/model"
	"WhaMan/pkg/global"
)

type CustomerImpl struct {
}

func (CustomerImpl) Create(customer *model.Customer) error {
	result := global.DB.Create(customer)
	return result.Error
}

func (CustomerImpl) Find(name string) (*model.Customer, error) {
	panic("implement me")
}

func (CustomerImpl) List() ([]*model.Customer, error) {
	panic("implement me")
}

func (CustomerImpl) ListSellOrders(name string) ([]*sellModel.SellOrder, error) {
	panic("implement me")
}

func (CustomerImpl) Update(customer *model.Customer) error {
	panic("implement me")
}

func (CustomerImpl) Delete(name string) error {
	panic("implement me")
}
