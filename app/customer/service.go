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
			return myErr.FieldDuplicate.SetDetail("客户名称已存在")
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
func (Service) Get(id uint) (*dto.GetRsp, error) {
	var data *dto.GetRsp
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&do.Customer{}).Where("id = ?", id).First(&data).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if err := tx.Model(&sellDO.SellOrder{}).
			Where("customer_id = ?", id).
			Order("date desc").
			Scan(&data.SellOrders).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		return nil
	})
	return data, err
}

// List 获取客户列表
func (Service) List(req *dto.ListReq) ([]*dto.ListRsp, error) {
	// 构造子查询
	sellOrderSubTx := database.DB.Model(&sellDO.SellOrder{}).Select("*")
	if req.SellOrdersWhere != nil {
		if req.SellOrdersWhere.Date != nil {
			if req.SellOrdersWhere.Date.StartDate != "" {
				sellOrderSubTx = sellOrderSubTx.Where("date >= ?", req.SellOrdersWhere.Date.StartDate)
			}
			if req.SellOrdersWhere.Date.EndDate != "" {
				sellOrderSubTx = sellOrderSubTx.Where("date < ?", req.SellOrdersWhere.Date.EndDate)
			}
		}
	}

	// 统计交易额、利润、已收款等信息
	tx := database.DB.Model(&do.Customer{}).
		Select("customers.*, SUM(so.unit_price*so.quantity) as turnover, SUM((so.unit_price-so.restock_unit_price)*so.quantity-so.freight_cost-so.kickback-so.tax-so.other_cost) as profit, SUM(so.paid_money) as paid_money ").
		Joins("LEFT JOIN (?) so ON customers.id = so.customer_id", sellOrderSubTx).
		Group("customers.id")
	if req.OrderBy == "" {
		req.OrderBy = "CONVERT(name USING gbk)" // 默认
	}
	tx = tx.Order(req.OrderBy)
	var customers []*dto.ListRsp
	if err := tx.Scan(&customers).Error; err != nil {
		log.Logger.Error(err)
		return nil, myErr.ServerErr
	}
	return customers, nil
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
				return myErr.FieldDuplicate.SetDetail("客户名称已存在")
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
		// 存在关联的出货订单，不能删除
		sellOrders := make([]sellDO.SellOrder, 0)
		if err := tx.Model(&sellDO.SellOrder{}).Where("customer_id = ?", id).
			Find(&sellOrders).Error; err != nil {
			log.Logger.Error(err)
			return myErr.ServerErr
		}
		if len(sellOrders) != 0 {
			return myErr.CannotDelete.SetDetail("存在关联的出货订单，不能删除")
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
