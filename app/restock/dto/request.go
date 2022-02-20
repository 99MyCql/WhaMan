package dto

import (
	"encoding/json"

	"WhaMan/app/restock/do"
	stockDO "WhaMan/app/stock/do"
	"WhaMan/pkg/datetime"
)

// ComReq Create,Update 接口请求参数
type ComReq struct {
	Date          datetime.MyDatetime `json:"date" binding:"required"`      // 日期
	ModelNum      string              `json:"model_num" binding:"required"` // 型号
	Specification string              `json:"specification"`                // 规格
	Quantity      float64             `json:"quantity"`                     // 数量
	UnitPrice     float64             `json:"unit_price"`                   // 单价
	SupplierID    uint                `json:"supplier_id"`                  // 供应商(外键)
	PaidMoney     float64             `json:"paid_money"`                   // 已付金额
	PayMethod     string              `json:"pay_method"`                   // 付款方式
	Note          string              `json:"note"`                         // 备注
	Location      string              `json:"location"`                     // 存放地点
}

// Convert2RestockOrder 根据进货信息生成进货订单
func (r *ComReq) Convert2RestockOrder() *do.RestockOrder {
	return &do.RestockOrder{
		Date:          r.Date,
		ModelNum:      r.ModelNum,
		Specification: r.Specification,
		Quantity:      r.Quantity,
		UnitPrice:     r.UnitPrice,
		SumMoney:      r.Quantity * r.UnitPrice,
		SupplierID:    r.SupplierID,
		PaidMoney:     r.PaidMoney,
		PayMethod:     r.PayMethod,
		Location:      r.Location,
		Note:          r.Note,
	}
}

func (r *ComReq) Convert2Stock() *stockDO.Stock {
	return &stockDO.Stock{
		ModelNum:        r.ModelNum,
		Specification:   r.Specification,
		RestockQuantity: r.Quantity,
		CurQuantity:     r.Quantity,
		UnitPrice:       r.UnitPrice,
		SumMoney:        r.Quantity * r.UnitPrice,
		Location:        r.Location,
		Note:            r.Note,
	}
}

type ListReq struct {
	Where *struct {
		Date *struct {
			StartDate string `binding:"datetime=2006-01-02"`
			EndDate   string `binding:"datetime=2006-01-02"`
		}
		SupplierID uint
		StockID    uint
	}
	OrderBy string
}

func (o *ListReq) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
