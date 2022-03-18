package dto

import (
	"encoding/json"

	"WhaMan/app/sell/do"
	"WhaMan/pkg/datetime"
)

type ComReq struct {
	Date            datetime.MyDatetime `json:"date" binding:"required"`       // 日期
	CustomerOrderID string              `json:"customer_order_id"`             // 客户订单号
	CustomerBatchID string              `json:"customer_batch_id"`             // 客户批号
	DeliverOrderID  string              `json:"deliver_order_id"`              // 送货单号
	ModelNum        string              `json:"model_num" binding:"required"`  // 型号
	Specification   string              `json:"specification"`                 // 规格
	Quantity        float64             `json:"quantity" binding:"required"`   // 数量
	UnitPrice       float64             `json:"unit_price" binding:"required"` // 单价
	PaidMoney       float64             `json:"paid_money"`                    // 已付金额（客户已经支付的金额）
	PayDate         datetime.MyDatetime `json:"pay_date"`                      // 付款日期
	PayMethod       string              `json:"pay_method"`                    // 付款方式
	FreightCost     float64             `json:"freight_cost"`                  // 运费
	Kickback        float64             `json:"kickback"`                      // 回扣
	Tax             float64             `json:"tax"`                           // 税金
	OtherCost       float64             `json:"other_cost"`                    // 杂费
	Note            string              `json:"note"`                          // 备注

	// 当 RestockOrderID 为 null 时，需给出进货单价
	RestockUnitPrice float64 `json:"restock_unit_price" gorm:"not null"` // 进货单价
	RestockOrderID   *uint   `json:"restock_order_id"`                   // 进货订单编号(外键)，可为null
	CustomerID       uint    `json:"customer_id" binding:"required"`     // 客户编号(外键)
}

func (r *ComReq) Convert2SellOrder() *do.SellOrder {
	return &do.SellOrder{
		Date:             r.Date,
		CustomerOrderID:  r.CustomerOrderID,
		CustomerBatchID:  r.CustomerBatchID,
		DeliverOrderID:   r.DeliverOrderID,
		ModelNum:         r.ModelNum,
		Specification:    r.Specification,
		Quantity:         r.Quantity,
		UnitPrice:        r.UnitPrice,
		PaidMoney:        r.PaidMoney,
		PayDate:          r.PayDate,
		PayMethod:        r.PayMethod,
		FreightCost:      r.FreightCost,
		Kickback:         r.Kickback,
		Tax:              r.Tax,
		OtherCost:        r.OtherCost,
		Note:             r.Note,
		RestockUnitPrice: r.RestockUnitPrice,
		RestockOrderID:   r.RestockOrderID,
		CustomerID:       r.CustomerID,
	}
}

// Date 左闭右开
type Date struct {
	StartDate string `json:"start_date" binding:"datetime=2006-01-02"`
	EndDate   string `json:"end_date" binding:"datetime=2006-01-02"`
}

type Where struct {
	Date           *Date `json:"date"`
	CustomerID     uint  `json:"customer_id"`
	RestockOrderID uint  `json:"restock_order_id"`
}

type ListReq struct {
	Where   *Where `json:"where"`
	OrderBy string `json:"order_by"`
}

func (r *ListReq) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
