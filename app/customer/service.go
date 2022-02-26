package customer

import (
	"WhaMan/app/customer/do"
	"WhaMan/app/customer/dto"
	sellDO "WhaMan/app/sell/do"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type Service struct{}

// Create 创建
func (s *Service) Create(req *dto.ComReq) (uint, error) {
	customer := req.Convert2Customer()
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查名称是否存在
		exist, err := s.nameIsExist(tx, customer.Name)
		if err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if exist {
			log.Logger.Info("客户名称已存在")
			return myErr.FieldDuplicate.AddMsg("客户名称已存在")
		}

		// 创建客户
		if err := tx.Create(customer).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return customer.ID, err
}

// Get 查找
func (Service) Get(id uint) (*dto.ComRsp, error) {
	var data *dto.ComRsp
	if err := database.DB.Model(&do.Customer{}).Where("id = ?", id).First(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	if err := database.DB.Model(&sellDO.SellOrder{}).Where("customer_id = ?", id).
		Scan(&data.SellOrders).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// List 获取客户列表
// TODO: 可根据条件进行筛选
func (Service) List() ([]*dto.ComRsp, error) {
	var data []*dto.ComRsp
	if err := database.DB.Model(&do.Customer{}).
		Order("CONVERT(name USING gbk)").
		Omit("SellOrders").Scan(&data).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return data, nil
}

// Update 更新
func (s *Service) Update(id uint, req *dto.ComReq) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查更新后的客户名称是否已经存在(未更新不检查)
		var oldCustomer *do.Customer
		if err := tx.First(&oldCustomer, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		newCustomer := req.Convert2Customer()
		if newCustomer.Name != oldCustomer.Name {
			exist, err := s.nameIsExist(tx, newCustomer.Name)
			if err != nil {
				log.Logger.Error(err)
				return myErr.ServerErr
			}
			if exist {
				log.Logger.Info("客户名称已存在")
				return myErr.FieldDuplicate.AddMsg("客户名称已存在")
			}
		}

		// 更新
		newCustomer.ID = id
		if err := tx.Select("*").Omit("Turnover", "UnpaidMoney", "CreatedAt").
			Updates(newCustomer).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// Delete 删除
func (Service) Delete(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查交易额，不为零不能删除
		var customer *do.Customer
		if err := tx.First(&customer, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return myErr.NotFound.AddMsg("客户ID不存在")
			}
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if customer.Turnover != 0 {
			return myErr.CannotDelete.AddMsg("交易额不为零不能删除")
		}

		// 删除
		if err := tx.Unscoped().Delete(&do.Customer{}, id).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
}

// nameIsExist 检查名称是否存在
func (Service) nameIsExist(tx *gorm.DB, name string) (bool, error) {
	if err := tx.Where("name = ?", name).First(&do.Customer{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrapf(err, "检查名称<%s>是否存在时出错", name)
	}
	return true, nil
}
