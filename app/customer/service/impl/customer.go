package impl

import (
	"WhaMan/app/customer/model"
	sellModel "WhaMan/app/sell/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
)

type CustomerImpl struct {
}

func (CustomerImpl) Create(i *model.CustomerInfo) error {
	if err := global.DB.Create(&model.Customer{
		Name:     i.Name,
		Contacts: i.Contacts,
		Phone:    i.Phone,
		Note:     i.Note,
	}).Error; err != nil {
		return errors.Wrapf(err, "创建客户失败：%+v", i)
	}
	return nil
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

func (CustomerImpl) Update(i *model.CustomerInfo) error {
	panic("implement me")
}

func (CustomerImpl) Delete(name string) error {
	panic("implement me")
}
