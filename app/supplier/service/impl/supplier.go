package impl

import (
	"WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
)

type Supplier struct{}

func (Supplier) Create(i *model.SupplierInfo) error {
	supplier := GenSupplier(i)
	if err := global.DB.Create(supplier).Error; err != nil {
		return errors.Wrapf(err, "创建供应商失败：%+v", i)
	}
	return nil
}

func (Supplier) Find(name string) (*model.Supplier, error) {
	var s model.Supplier
	res := global.DB.Find(&s, name)
	if res.Error != nil {
		return nil, errors.Wrapf(res.Error, "查询供应商失败，供应商名：%s", name)
	}
	return &s, nil
}

func (Supplier) List() ([]*model.Supplier, error) {
	panic("implement me")
}

func (Supplier) Update(i *model.SupplierInfo) error {
	s := GenSupplier(i)
	res := global.DB.Model(s).Updates(s)
	if res.Error != nil {
		return errors.Wrapf(res.Error, "更新供应商信息失败，供应商信息：%+v", *i)
	}
	return nil
}

func (Supplier) Delete(name string) error {
	panic("implement me")
}

// GenSupplier 根据 SupplierInfo 生成 Supplier
func GenSupplier(i *model.SupplierInfo) *model.Supplier {
	return &model.Supplier{
		Name:     i.Name,
		Contacts: i.Contacts,
		Phone:    i.Phone,
		Note:     i.Note,
	}
}
