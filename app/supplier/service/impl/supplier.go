package impl

import (
	"WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type SupplierImpl struct{}

// Create 1.检查名称是否存在；2.创建
func (s *SupplierImpl) Create(p *model.Params) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		newSupplier := p.GenSupplier()

		// 检查名称是否已经存在
		if err := s.checkName(tx, newSupplier.Name); err != nil {
			return errors.WithMessagef(err, "创建供应商过程中，检查名称是否已存在出错：%+v", newSupplier)
		}

		if err := tx.Create(newSupplier).Error; err != nil {
			return errors.Wrapf(err, "创建供应商失败：%+v", newSupplier)
		}
		return nil
	})
}

// Find 查找
func (SupplierImpl) Find(id uint) (*model.Supplier, error) {
	var supplier *model.Supplier
	if err := global.DB.First(&supplier, id).Error; err != nil {
		return nil, errors.Wrapf(err, "通过ID查询供应商出错：%d", id)
	}
	return supplier, nil
}

// List 获取列表
func (SupplierImpl) List() ([]*model.Supplier, error) {
	var suppliers []*model.Supplier
	if err := global.DB.Find(&suppliers).Error; err != nil {
		return nil, errors.Wrapf(err, "查询供应商列表出错")
	}
	return suppliers, nil
}

// Update 1.检查更新后的名称是否已经存在(未更新不检查)；2.更新
func (s *SupplierImpl) Update(id uint, p *model.Params) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var oldSupplier *model.Supplier
		if err := tx.First(&oldSupplier, id).Error; err != nil {
			return errors.Wrapf(err, "更新供应商信息过程中，通过ID查询供应商时出错：%d", id)
		}
		newSupplier := p.GenSupplier()
		newSupplier.Model = oldSupplier.Model

		// 检查客户名是否已经存在
		if newSupplier.Name != oldSupplier.Name {
			if err := s.checkName(tx, newSupplier.Name); err != nil {
				return errors.WithMessagef(err, "更新供应商信息过程中，检查名称是否已存在出错：%+v", newSupplier)
			}
		}

		if err := tx.Select("*").Omit("Turnover").Updates(newSupplier).Error; err != nil {
			return errors.Wrapf(err, "更新供应商信息失败：%+v", newSupplier)
		}
		return nil
	})
}

// Delete 1.交易额不为零不能删除；2.删除
func (s *SupplierImpl) Delete(id uint) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var supplier *model.Supplier
		if err := tx.First(&supplier, id).Error; err != nil {
			return errors.Wrapf(err, "删除供应商过程中，通过ID查询供应商出错：%d", id)
		}
		if supplier.Turnover != 0 {
			return errors.WithMessagef(global.ErrCannotDelete, "删除供应商过程中，交易额不为零不能删除：%+v", supplier)
		}

		if err := tx.Unscoped().Delete(&model.Supplier{}, id).Error; err != nil {
			return errors.Wrapf(err, "删除供应商失败：%+v", supplier)
		}
		return nil
	})
}

// checkName 检查名称是否存在
func (s SupplierImpl) checkName(tx *gorm.DB, name string) error {
	err := tx.Where("name = ?", name).First(&model.Supplier{}).Error
	if err == nil {
		// 未返回错误，说明通过名称查询到了数据，进而说明名称已存在
		return errors.WithMessagef(global.ErrNameExist, "名称已存在：%s", name)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 返回错误，但错误不是“记录未找到”，则说明查询过程中出现了其它错误
		return errors.Wrapf(err, "根据名称查询供应商出错：%s", name)
	}
	return nil
}
