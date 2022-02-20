package stock

import (
	sellDO "WhaMan/app/sell/do"
	"WhaMan/app/stock/do"
	"WhaMan/app/stock/dto"
	"WhaMan/pkg/database"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type Service struct{}

// Get 查找
func (Service) Get(id uint) (*dto.ComRsp, error) {
	data := &dto.ComRsp{}
	err := database.DB.Model(&do.Stock{}).
		Select("stocks.*, restock_orders.date as restock_date, suppliers.name as supplier_name").
		Joins("JOIN restock_orders ON restock_orders.stock_id = stocks.id").
		Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id ").
		Where("stocks.id = ?", id).
		Scan(&data).Error
	if err != nil {
		return nil, myErr.ServerErr
	}
	log.Logger.Debugf("%+v", data)
	return data, nil
}

// List 获取列表
func (Service) List(req *dto.ListReq) ([]*dto.ComRsp, error) {
	data := make([]*dto.ComRsp, 0)
	tx := database.DB.Model(&do.Stock{})

	if req.Where != nil {
		if req.Where.CurQuantity != nil {
			if req.Where.CurQuantity.Start != nil {
				tx = tx.Where("cur_quantity >= ?", *req.Where.CurQuantity.Start)
			}
			if req.Where.CurQuantity.End != nil {
				tx = tx.Where("cur_quantity <= ?", *req.Where.CurQuantity.End)
			}
		}
	}
	if req.OrderBy != "" {
		tx = tx.Order(req.OrderBy)
	}
	err := tx.Select("stocks.*, restock_orders.date as restock_date, suppliers.name as supplier_name").
		Joins("JOIN restock_orders ON restock_orders.stock_id = stocks.id").
		Joins("JOIN suppliers ON restock_orders.supplier_id = suppliers.id ").
		Scan(&data).Error
	if err != nil {
		log.Logger.Error(err)
		return data, myErr.ServerErr
	}
	return data, nil
}

// Update
// 仅更新备注、存放地点：1.更新
// 更新数据较多时：1.更新对应的进货订单；2.更新对应的出货订单(单价)；3.重新计算当前库存和金额；4.更新库存
func (Service) Update(id uint, req *dto.UpdateReq) error {
	if err := database.DB.Where("id = ?", id).Select("location", "note").
		Updates(req.Convert2Stock()).Error; err != nil {
		log.Logger.Error(err)
		return myErr.ServerErr
	}
	return nil
}

// UpdateSellOrders 当单价发生变化时，更新关联出货订单的利润
func (Service) UpdateSellOrders(tx *gorm.DB, id uint, unitPrice float64) error {
	var sellOrders []*sellDO.SellOrder
	if err := tx.Where("stock_id = ?", id).Find(&sellOrders).Error; err != nil {
		return errors.Wrapf(err, "更新关联的出货订单过程中，查询出货订单出错：%d", id)
	}
	for i := 0; i < len(sellOrders); i++ {
		sellOrders[i].RestockUnitPrice = unitPrice
		sellOrders[i].CalProfit()
		if err := tx.Model(&sellOrders[i]).Update("profit", sellOrders[i].Profit).Error; err != nil {
			return errors.Wrapf(err, "更新关联的出货订单过程中，更新出货订单出错：%+v", sellOrders[i])
		}
	}
	return nil
}
