package impl

import (
	"fmt"

	restockModel "WhaMan/app/restock/model"
	stockModel "WhaMan/app/stock/model"
	supplierModel "WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Restock struct {
}

// Restock 新增进货订单，并新增库存，以及更新供应商交易额信息
func (r *Restock) Restock(i *restockModel.RestockInfo) error {
	// 生成进货订单ID=库存ID
	var id string
	for index := 1; ; index++ {
		id = GenID(i, index)
		if id == "" {
			return errors.Errorf("进货流程中，ID超出编号范围：%s", id)
		}
		if err := global.DB.First(&restockModel.RestockOrder{}, "id = ?", id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			} else {
				return errors.Wrapf(err, "进货流程中，查询进货订单ID是否重复失败：%s", id)
			}
		}
	}

	// 执行事务：保证新增库存、新增进货订单、更新供应商信息等操作同时成功
	restockOrder := GenRestockOrder(i, id)
	stock := GenStock(i, id)
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(stock).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，新增库存失败：%+v", stock)
		}
		restockOrder.StockID = stock.ID
		if err := tx.Create(restockOrder).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，新增进货订单失败：%+v", restockOrder)
		}
		var supplier supplierModel.Supplier
		if err := tx.First(&supplier, "name = ?", i.SupplierName).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，获取供应商信息失败：%s", i.SupplierName)
		}
		supplier.Turnover += i.SumMoney
		if err := tx.Save(supplier).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，更新供应商信息失败：%+v", supplier)
		}
		return nil
	})
}

func (r *Restock) Find(id string) (*restockModel.RestockOrder, error) {
	panic("implement me")
}

func (r *Restock) List() ([]*restockModel.RestockOrder, error) {
	panic("implement me")
}

func (r *Restock) Update(i *restockModel.RestockInfo) error {
	panic("implement me")
}

func (r *Restock) Delete(id string) error {
	panic("implement me")
}

// GenRestockOrder 根据进货信息生成进货订单
func GenRestockOrder(i *restockModel.RestockInfo, id string) *restockModel.RestockOrder {
	return &restockModel.RestockOrder{
		ID:            id,
		Date:          i.Date,
		Model:         i.Model,
		Specification: i.Specification,
		Quantity:      i.Quantity,
		Unit:          i.Unit,
		UnitPrice:     i.UnitPrice,
		SumMoney:      i.SumMoney,
		SupplierName:  i.SupplierName,
		PaidMoney:     i.PaidMoney,
		PayMethod:     i.PayMethod,
		Note:          i.Note,
	}
}

// GenStock 根据进货信息生成库存信息
func GenStock(i *restockModel.RestockInfo, id string) *stockModel.Stock {
	return &stockModel.Stock{
		ID:              id,
		Model:           i.Model,
		Specification:   i.Specification,
		RestockQuantity: i.Quantity,
		SellQuantity:    0,
		CurQuantity:     i.Quantity,
		Unit:            i.Unit,
		UnitPrice:       i.UnitPrice,
		SumMoney:        i.SumMoney,
		Location:        i.Location,
	}
}

func GenID(i *restockModel.RestockInfo, index int) string {
	if index >= 10000 {
		return ""
	}
	return i.Date.Format("20060102") + "-" + i.SupplierName + "-" + i.Model + fmt.Sprintf("-%04d", index)
}
