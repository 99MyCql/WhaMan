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
	// 检查名称是否已经存在
	_, err := s.FindByName(p.Name)
	if err == nil {
		// 未返回错误，说明通过名称查询到了数据，进而说明名称已存在
		return errors.WithMessagef(global.ErrNameExist, "创建供应商过程中，名称已存在：%+v", p)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 返回错误，但错误不是“记录未找到”，则说明查询过程中出现了其它错误
		return errors.Wrapf(err, "创建供应商过程中，检查名称是否已存在出错：%+v", p)
	}

	if err := global.DB.Create(p.GenSupplier()).Error; err != nil {
		return errors.Wrapf(err, "创建供应商失败：%+v", p)
	}
	return nil
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
	// 检查客户名是否已经存在
	anotherSupplier, err := s.FindByName(p.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询过程报错，且错误不是记录未找到，则说明查询过程中出现了其它错误
		return errors.WithMessagef(err, "更新供应商信息过程中，检查名称是否已存在时出错：%d-%+v", id, p)
	} else if err == nil && anotherSupplier.ID != id {
		// 查询未报错，但查询到的供应商ID不是当前更新ID，说明更新后的供应商名称已存在/出现重复
		return errors.WithMessagef(global.ErrNameExist, "更新供应商信息过程中，名称已存在：%d-%+v", id, p)
	}

	if err := global.DB.Where("id = ?", id).Updates(p.GenSupplier()).Error; err != nil {
		return errors.Wrapf(err, "更新供应商信息失败：%d-%+v", id, p)
	}
	return nil
}

// Delete 1.交易额不为零不能删除；2.删除
func (s *SupplierImpl) Delete(id uint) error {
	// 交易额不为零不能删除
	supplier, err := s.Find(id)
	if err != nil {
		return errors.WithMessagef(err, "删除供应商过程中，根据ID查询出错：%d", id)
	}
	if supplier.Turnover != 0 {
		return errors.WithMessagef(global.ErrCannotDelete, "删除供应商过程中，交易额不为零不能删除：%d", id)
	}

	if err := global.DB.Unscoped().Delete(&model.Supplier{}, id).Error; err != nil {
		return errors.Wrapf(err, "删除供应商失败：%d", id)
	}
	return nil
}

func (SupplierImpl) FindByName(name string) (*model.Supplier, error) {
	var supplier *model.Supplier
	if err := global.DB.Where("name = ?", name).First(&supplier).Error; err != nil {
		return nil, errors.Wrapf(err, "通过名称查询供应商时出错：%s", name)
	}
	return supplier, nil
}
