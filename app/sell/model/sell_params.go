package model

import (
	"WhaMan/pkg/global/models"
)

type SellParams struct {
	Date             models.MyDatetime `json:"date" binding:"required"`        // 日期
	CustomerOrderID  string            `json:"customer_order_id"`              // 客户订单号
	CustomerBatchID  string            `json:"customer_batch_id"`              // 客户批号
	DeliverOrderID   string            `json:"deliver_order_id"`               // 送货单号
	ModelNum         string            `json:"model_num" binding:"required"`   // 型号
	Specification    string            `json:"specification"`                  // 规格
	Quantity         float64           `json:"quantity"`                       // 数量
	RestockUnitPrice float64           `json:"restock_unit_price"`             // 进货单价
	UnitPrice        float64           `json:"unit_price"`                     // 单价
	PaidMoney        float64           `json:"paid_money"`                     // 已付金额
	PayDate          models.MyDatetime `json:"pay_date"`                       // 付款日期
	PayMethod        string            `json:"pay_method"`                     // 付款方式
	FreightCost      float64           `json:"freight_cost"`                   // 运费
	Kickback         float64           `json:"kickback"`                       // 回扣
	Tax              float64           `json:"tax"`                            // 税金
	OtherCost        float64           `json:"other_cost"`                     // 杂费
	Note             string            `json:"note"`                           // 备注
	StockID          uint              `json:"stock_id"`                       // 库存编号(外键)
	CustomerID       uint              `json:"customer_id" binding:"required"` // 客户编号(外键)
}

func (p *SellParams) GenSellOrder() *SellOrder {
	return &SellOrder{
		Date:             p.Date,
		CustomerOrderID:  p.CustomerOrderID,
		CustomerBatchID:  p.CustomerBatchID,
		DeliverOrderID:   p.DeliverOrderID,
		ModelNum:         p.ModelNum,
		Specification:    p.Specification,
		Quantity:         p.Quantity,
		RestockUnitPrice: p.RestockUnitPrice,
		UnitPrice:        p.UnitPrice,
		SumMoney:         p.Quantity * p.UnitPrice,
		PaidMoney:        p.PaidMoney,
		PayDate:          p.PayDate,
		PayMethod:        p.PayMethod,
		FreightCost:      p.FreightCost,
		Kickback:         p.Kickback,
		Tax:              p.Tax,
		OtherCost:        p.OtherCost,
		Profit:           0,
		Note:             p.Note,
		StockID:          p.StockID,
		CustomerID:       p.CustomerID,
	}
}
