package dto

import (
	"encoding/json"

	"WhaMan/app/sell/do"
	"WhaMan/pkg/datetime"
)

type ComReq struct {
	Date             datetime.MyDatetime `json:"date" binding:"required"`        // 日期
	CustomerOrderID  string              `json:"customer_order_id"`              // 客户订单号
	CustomerBatchID  string              `json:"customer_batch_id"`              // 客户批号
	DeliverOrderID   string              `json:"deliver_order_id"`               // 送货单号
	ModelNum         string              `json:"model_num" binding:"required"`   // 型号
	Specification    string              `json:"specification"`                  // 规格
	Quantity         float64             `json:"quantity"`                       // 数量
	RestockUnitPrice float64             `json:"restock_unit_price"`             // 进货单价
	UnitPrice        float64             `json:"unit_price"`                     // 单价
	PaidMoney        float64             `json:"paid_money"`                     // 已付金额（客户已经支付的金额）
	PayDate          datetime.MyDatetime `json:"pay_date"`                       // 付款日期
	PayMethod        string              `json:"pay_method"`                     // 付款方式
	FreightCost      float64             `json:"freight_cost"`                   // 运费
	Kickback         float64             `json:"kickback"`                       // 回扣
	Tax              float64             `json:"tax"`                            // 税金
	OtherCost        float64             `json:"other_cost"`                     // 杂费
	Note             string              `json:"note"`                           // 备注
	StockID          *uint               `json:"stock_id"`                       // 库存编号(外键)，为null则不关联库存
	CustomerID       uint                `json:"customer_id" binding:"required"` // 客户编号(外键)
}

func (r *ComReq) Convert2SellOrder() *do.SellOrder {
	s := &do.SellOrder{
		Date:             r.Date,
		CustomerOrderID:  r.CustomerOrderID,
		CustomerBatchID:  r.CustomerBatchID,
		DeliverOrderID:   r.DeliverOrderID,
		ModelNum:         r.ModelNum,
		Specification:    r.Specification,
		Quantity:         r.Quantity,
		RestockUnitPrice: r.RestockUnitPrice,
		UnitPrice:        r.UnitPrice,
		SumMoney:         r.Quantity * r.UnitPrice,
		PaidMoney:        r.PaidMoney,
		PayDate:          r.PayDate,
		PayMethod:        r.PayMethod,
		FreightCost:      r.FreightCost,
		Kickback:         r.Kickback,
		Tax:              r.Tax,
		OtherCost:        r.OtherCost,
		Note:             r.Note,
		StockID:          r.StockID,
		CustomerID:       r.CustomerID,
	}
	s.CalProfit()
	return s
}

type ListReq struct {
	Where *struct {
		Date *struct {
			StartDate string `json:"start_date" binding:"datetime=2006-01-02"`
			EndDate   string `json:"end_date" binding:"datetime=2006-01-02"`
		} `json:"date"`
		CustomerID uint `json:"customer_id"`
		StockID    uint `json:"stock_id"`
	} `json:"where"`
	OrderBy string `json:"order_by"`
}

func (r *ListReq) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
