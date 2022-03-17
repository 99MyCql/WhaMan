package dto

import (
	"encoding/json"

	"WhaMan/app/restock/do"
	"WhaMan/pkg/datetime"
)

// ComReq Create Update 接口请求参数
type ComReq struct {
	Date          datetime.MyDatetime `json:"date" binding:"required"`      // 日期
	ModelNum      string              `json:"model_num" binding:"required"` // 型号
	Specification string              `json:"specification"`                // 规格
	Quantity      float64             `json:"quantity"`                     // 数量
	UnitPrice     float64             `json:"unit_price"`                   // 单价
	PaidMoney     float64             `json:"paid_money"`                   // 已付金额
	PayMethod     string              `json:"pay_method"`                   // 付款方式
	Note          string              `json:"note"`                         // 备注
	Location      string              `json:"location"`                     // 存放地点
	SupplierID    uint                `json:"supplier_id"`                  // 供应商编号(外键)
}

// Convert2RestockOrder 根据进货信息生成进货订单
func (r *ComReq) Convert2RestockOrder() *do.RestockOrder {
	return &do.RestockOrder{
		Date:          r.Date,
		ModelNum:      r.ModelNum,
		Specification: r.Specification,
		Quantity:      r.Quantity,
		UnitPrice:     r.UnitPrice,
		PaidMoney:     r.PaidMoney,
		PayMethod:     r.PayMethod,
		Location:      r.Location,
		Note:          r.Note,
		SupplierID:    r.SupplierID,
	}
}

type Date struct {
	StartDate string `json:"start_date" binding:"datetime=2006-01-02"`
	EndDate   string `json:"end_date" binding:"datetime=2006-01-02"`
}

type Where struct {
	Date       *Date  `json:"date"`
	SupplierID uint   `json:"supplier_id"`
	ModelNum   string `json:"model_num"`
}

type SellOrdersWhere struct {
	Date *Date `json:"date"`
}

type ListReq struct {
	Where           *Where           `json:"where"`
	OrderBy         string           `json:"order_by"`
	WithSellOrders  bool             `json:"with_sell_orders"`
	SellOrdersWhere *SellOrdersWhere `json:"sell_orders_where"`
}

func (o *ListReq) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

type ListGroupByModelNumReq struct {
	OrderBy string `json:"order_by"`
}

func (o *ListGroupByModelNumReq) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
