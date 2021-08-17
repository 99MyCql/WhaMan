package impl

import (
	"WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type Supplier struct{}

// Create 1.检查名称是否存在；2.创建
func (s *Supplier) Create(p *model.Params) error {
	// 检查名称是否已经存在
	_, err := s.FindByName(p.Name)
	if err == nil {
		// 未返回错误，说明通过名称查询到了数据，进而说明名称已存在
		return global.ErrNameExist
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 返回错误，但错误不是“记录未找到”，则说明查询过程中出现了其它错误
		return err
	}

	if err := global.DB.Create(p.GenSupplier()).Error; err != nil {
		return errors.Wrapf(err, "创建供应商失败：%+v", p)
	}
	return nil
}

// Find 查找
func (Supplier) Find(id uint) (*model.Supplier, error) {
	var supplier *model.Supplier
	if err := global.DB.First(&supplier, id).Error; err != nil {
		return nil, errors.Wrapf(err, "通过ID查询供应商出错：%d", id)
	}
	return supplier, nil
}

// List 获取列表
func (Supplier) List() ([]*model.Supplier, error) {
	var suppliers []*model.Supplier
	if err := global.DB.Find(&suppliers).Error; err != nil {
		return nil, errors.Wrapf(err, "查询供应商列表出错")
	}
	return suppliers, nil
}

// Update 1.检查更新权限；2.检查更新后的名称是否已经存在(未更新不检查)；3.更新
func (s *Supplier) Update(id uint, p *model.Params) error {
	// TODO: 检查是否具有更新该id对应记录的权限

	// 检查客户名是否已经存在
	anotherSupplier, err := s.FindByName(p.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询过程报错，且错误不是记录未找到，则说明查询过程中出现了其它错误
		return err
	} else if err == nil && anotherSupplier.ID != id {
		// 查询未报错，但查询到的供应商ID不是当前更新ID，说明更新后的供应商名称已存在/出现重复
		return global.ErrNameExist
	}

	if err := global.DB.Where("id = ?", id).Updates(p.GenSupplier()).Error; err != nil {
		return errors.Wrapf(err, "更新供应商信息失败：%d-%+v", id, p)
	}
	return nil
}

// Delete 1.检查删除权限；2.删除
func (Supplier) Delete(id uint) error {
	// TODO: 检查是否具有删除该id对应记录的权限

	if err := global.DB.Unscoped().Delete(&model.Supplier{}, id).Error; err != nil {
		return errors.Wrapf(err, "删除供应商失败：%d", id)
	}
	return nil
}

func (Supplier) FindByName(name string) (*model.Supplier, error) {
	var supplier *model.Supplier
	if err := global.DB.Where("name = ?", name).First(&supplier).Error; err != nil {
		return nil, errors.Wrapf(err, "通过名称查询供应商时出错：%s", name)
	}
	return supplier, nil
}
