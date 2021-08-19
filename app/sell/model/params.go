package model

type Params struct {
	Date            string  `binding:"required,datetime=2006-01-02"` // 日期
	CustomerOrderID string  // 客户订单号
	CustomerBatchID string  // 客户批号
	DeliverOrderID  string  // 送货单号
	Specification   string  // 规格
	Quantity        float64 // 数量
	UnitPrice       float64 // 单价
	PaidMoney       float64 // 已付金额
	PayMethod       string  // 付款方式
	FreightCost     float64 // 运费
	Kickback        float64 // 回扣
	Tax             float64 // 税金
	OtherCost       float64 // 杂费
	Note            string  // 备注
	StockID         uint    `binding:"required"` // 库存编号(外键)
	CustomerID      uint    `binding:"required"` // 客户编号(外键)
}

func (p *Params) GenSellOrder() *SellOrder {
	return &SellOrder{
		Date:            p.Date,
		CustomerOrderID: p.CustomerOrderID,
		CustomerBatchID: p.CustomerBatchID,
		DeliverOrderID:  p.DeliverOrderID,
		Specification:   p.Specification,
		Quantity:        p.Quantity,
		UnitPrice:       p.UnitPrice,
		SumMoney:        p.Quantity * p.UnitPrice,
		PaidMoney:       p.PaidMoney,
		PayMethod:       p.PayMethod,
		FreightCost:     p.FreightCost,
		Kickback:        p.Kickback,
		Tax:             p.Tax,
		OtherCost:       p.OtherCost,
		Profit:          0,
		Note:            p.Note,
		StockID:         p.StockID,
		CustomerID:      p.CustomerID,
	}
}
