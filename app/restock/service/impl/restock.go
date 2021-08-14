package impl

import (
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
	// TODO: ID重复
	restockOrder := GenRestockOrder(i)
	stock := GenStock(i)
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(stock).Error; err != nil {
			return errors.Wrapf(err, "新增库存失败：%+v", stock)
		}
		restockOrder.StockID = stock.ID
		if err := tx.Create(restockOrder).Error; err != nil {
			return errors.Wrapf(err, "新增进货订单失败：%+v", restockOrder)
		}
		var supplier supplierModel.Supplier
		if err := tx.Where("name = ?", i.SupplierName).First(&supplier).Error; err != nil {
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
func GenRestockOrder(i *restockModel.RestockInfo) *restockModel.RestockOrder {
	return &restockModel.RestockOrder{
		ID:            GenID(i),
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
func GenStock(i *restockModel.RestockInfo) *stockModel.Stock {
	return &stockModel.Stock{
		ID:              GenID(i),
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

func GenID(i *restockModel.RestockInfo) string {
	return i.Date.Format("20060102") + "-" + i.SupplierName + "-" + i.Model + "-0001"
}
