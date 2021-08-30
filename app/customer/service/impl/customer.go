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
func (c *CustomerImpl) Create(p *model.Params) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		newCustomer := p.GenCustomer()

		if err := c.checkName(tx, newCustomer.Name); err != nil {
			return errors.WithMessagef(err, "创建客户过程中，检查名称是否已存在时出错：%+v", newCustomer)
		}

		if err := tx.Create(newCustomer).Error; err != nil {
			return errors.Wrapf(err, "创建客户失败：%+v", newCustomer)
		}
		return nil
	})
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
	if err := global.DB.Order("name").Find(&customers).Error; err != nil {
		return nil, errors.Wrapf(err, "查询客户列表出错")
	}
	return customers, nil
}

// Update 1.检查更新后的客户名称是否已经存在(未更新不检查)；2.更新
func (c *CustomerImpl) Update(id uint, p *model.Params) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var oldCustomer *model.Customer
		if err := tx.First(&oldCustomer, id).Error; err != nil {
			return errors.Wrapf(err, "更新客户信息过程中，通过ID查询客户出错：%d", id)
		}
		newCustomer := p.GenCustomer()
		newCustomer.Model = oldCustomer.Model

		// 如果客户名字变更，检查客户名是否已经存在
		if newCustomer.Name != oldCustomer.Name {
			if err := c.checkName(tx, newCustomer.Name); err != nil {
				return errors.WithMessagef(err, "更新客户信息过程中，检查名称是否已存在时出错：%+v", newCustomer)
			}
		}

		// 更新
		if err := tx.Select("*").Omit("Turnover", "UnpaidMoney").Updates(newCustomer).Error; err != nil {
			return errors.Wrapf(err, "更新客户信息失败：%+v", newCustomer)
		}
		return nil
	})
}

// Delete 1.交易额不为零不能删除；2.删除
func (c *CustomerImpl) Delete(id uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var customer *model.Customer
		if err := tx.First(&customer, id).Error; err != nil {
			return errors.Wrapf(err, "删除客户过程中，通过ID查询客户出错：%d", id)
		}

		if customer.Turnover != 0 {
			return errors.WithMessagef(global.ErrCannotDelete, "删除客户过程中，交易额不为零不能删除：%d", id)
		}

		if err := tx.Unscoped().Delete(&model.Customer{}, id).Error; err != nil {
			return errors.Wrapf(err, "删除客户失败：%d", id)
		}
		return nil
	})
}

// checkName 检查名称是否存在
func (c CustomerImpl) checkName(tx *gorm.DB, name string) error {
	err := tx.Where("name = ?", name).First(&model.Customer{}).Error
	if err == nil {
		return errors.WithMessagef(global.ErrNameExist, "名称已存在：%s", name)
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(err, "根据名称查询客户出错：%s", name)
	}
	return nil
}
