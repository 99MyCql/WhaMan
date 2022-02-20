package supplier

import (
	"WhaMan/app/supplier/do"
	"WhaMan/app/supplier/dto"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type Service struct{}

// Create 创建
func (s *Service) Create(p *dto.ComReq) (uint, error) {
	supplier := p.Convert2Supplier()
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查名称是否已经存在
		exist, err := s.nameIsExist(tx, supplier.Name)
		if err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if exist {
			log.Logger.Info("供应商名称已存在")
			return myErr.FieldDuplicate.AddMsg("供应商名称已存在")
		}

		// 创建
		if err := tx.Create(supplier).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return supplier.ID, err
}

// Get 查找
func (Service) Get(id uint) (*dto.ComRsp, error) {
	var data *dto.ComRsp
	if err := database.DB.Model(&do.Supplier{}).Where("id = ?", id).First(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// List 获取列表
func (Service) List() ([]*dto.ComRsp, error) {
	var suppliers []*dto.ComRsp
	if err := database.DB.Model(&do.Supplier{}).
		Order("CONVERT(name USING gbk)").Scan(&suppliers).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return suppliers, nil
}

// Update 更新
func (s *Service) Update(id uint, p *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查名称是否已经存在(名称未更新不检查)
		var oldSupplier *do.Supplier
		if err := tx.First(&oldSupplier, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		newSupplier := p.Convert2Supplier()
		if newSupplier.Name != oldSupplier.Name {
			exist, err := s.nameIsExist(tx, newSupplier.Name)
			if err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
			if exist {
				log.Logger.Info("供应商名称已存在")
				return myErr.FieldDuplicate.AddMsg("供应商名称已存在")
			}
		}

		// 更新
		newSupplier.ID = id
		if err := tx.Select("*").Omit("Turnover", "CreatedAt").
			Updates(newSupplier).Error; err != nil {
			return errors.Wrapf(err, "更新供应商信息失败：%+v", newSupplier)
		}

		return nil
	})
}

// Delete 删除
func (s *Service) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 交易额不为零不能删除
		var supplier *do.Supplier
		if err := tx.First(&supplier, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return myErr.NotFound.AddMsg("供应商ID不存在")
			}
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if supplier.Turnover != 0 {
			return myErr.CannotDelete.AddMsg("交易额不为零不能删除")
		}

		if err := tx.Unscoped().Delete(&do.Supplier{}, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// nameIsExist 检查名称是否存在
func (Service) nameIsExist(tx *gorm.DB, name string) (bool, error) {
	err := tx.Where("name = ?", name).First(&do.Supplier{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrapf(err, "检查名称<%s>是否存在时出错", name)
	}
	return true, nil
}
