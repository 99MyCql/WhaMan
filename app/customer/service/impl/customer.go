package impl

import (
	"WhaMan/app/customer/model"
	"WhaMan/pkg/global"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type CustomerImpl struct {
}

// Create 1.检查名称是否存在；2.创建客户
func (c *CustomerImpl) Create(customer *model.Customer) error {
	// 检查客户名称是否已经存在
	_, err := c.FindByName(customer.Name)
	if err == nil {
		// 未返回错误，说明通过名称查询到客户，进而说明名称已存在
		return global.ErrNameExist
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 返回错误，但错误不是记录未找到，则说明查询过程中出现了其它错误
		return err
	}

	if err := global.DB.Create(customer).Error; err != nil {
		return errors.Wrapf(err, "创建客户失败：%+v", customer)
	}
	return nil
}

// Find 查找
func (CustomerImpl) Find(id uint) (*model.Customer, error) {
	var customer *model.Customer
	if err := global.DB.First(&customer, id).Error; err != nil {
		return nil, errors.Wrapf(err, "通过ID查询客户出错：%d", id)
	}
	return customer, nil
}

// List 获取客户列表
func (CustomerImpl) List() ([]*model.Customer, error) {
	var customers []*model.Customer
	if err := global.DB.Find(&customers).Error; err != nil {
		return nil, errors.Wrapf(err, "查询客户列表出错")
	}
	return customers, nil
}

// Update 1.检查更新权限；2.检查更新后的客户是否已经存在(未更新不检查)；3.更新
func (c *CustomerImpl) Update(id uint, customer *model.Customer) error {
	// TODO: 检查是否具有更新该id对应记录的权限

	// 检查客户名是否已经存在
	anotherCustomer, err := c.FindByName(customer.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询过程报错，且错误不是记录未找到，则说明查询过程中出现了其它错误
		return err
	} else if err == nil && anotherCustomer.ID != id {
		// 查询未报错，但查询到的客户ID不是当前更新ID，说明更新后的客户名称已存在/出现重复
		return global.ErrNameExist
	}

	if err := global.DB.Where("id = ?", id).Updates(customer).Error; err != nil {
		return errors.Wrapf(err, "更新客户信息失败：%d-%+v", id, customer)
	}
	return nil
}

// Delete 1.检查删除权限；2.删除
func (CustomerImpl) Delete(id uint) error {
	// TODO: 检查是否具有删除该id对应记录的权限

	if err := global.DB.Delete(&model.Customer{}, id).Error; err != nil {
		return errors.Wrapf(err, "删除客户失败：%d", id)
	}
	return nil
}

func (CustomerImpl) FindByName(name string) (*model.Customer, error) {
	var customer *model.Customer
	if err := global.DB.Where("name = ?", name).First(&customer).Error; err != nil {
		return nil, errors.Wrapf(err, "通过客户名查询客户时出错：%s", name)
	}
	return customer, nil
}
