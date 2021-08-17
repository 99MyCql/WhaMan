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
	stock := GenStock(i)
	stock.RestockOrder = GenRestockOrder(i)
	// 获取关联的供应商
	var supplier supplierModel.Supplier
	if err := global.DB.First(&supplier, "id = ?", i.SupplierID).Error; err != nil {
		return errors.Wrapf(err, "进货流程中，获取供应商信息失败：%d", i.SupplierID)
	}
	supplier.Turnover += i.SumMoney
	// 执行事务：保证新增库存和进货订单、更新供应商信息等操作同时成功
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(stock).Error; err != nil {
			return errors.Wrapf(err, "进货流程中，新增库存和进货订单失败：%+v", stock)
		}
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
		Date:          i.Date,
		ModelNum:      i.ModelNum,
		Specification: i.Specification,
		Quantity:      i.Quantity,
		Unit:          i.Unit,
		UnitPrice:     i.UnitPrice,
		SumMoney:      i.SumMoney,
		SupplierID:    i.SupplierID,
		PaidMoney:     i.PaidMoney,
		PayMethod:     i.PayMethod,
		Note:          i.Note,
	}
}

// GenStock 根据进货信息生成库存信息
func GenStock(i *restockModel.RestockInfo) *stockModel.Stock {
	return &stockModel.Stock{
		ModelNum:        i.ModelNum,
		Specification:   i.Specification,
		RestockQuantity: i.Quantity,
		SellQuantity:    0,
		CurQuantity:     i.Quantity,
		UnitPrice:       i.UnitPrice,
		SumMoney:        i.SumMoney,
		Location:        i.Location,
	}
}

// func GenID(i *restockModel.RestockInfo, index int) string {
// 	if index >= 10000 {
// 		return ""
// 	}
// 	return i.Date.Format("20060102") + "-" + i.SupplierName + "-" + i.Model + fmt.Sprintf("-%04d", index)
// }
