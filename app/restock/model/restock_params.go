package model

// RestockParams 进货接口参数
type RestockParams struct {
	Date          string  `json:"date" binding:"required,datetime=2006-01-02"` // 日期(字符串形式)
	ModelNum      string  `json:"model_num" binding:"required"`                // 型号
	Specification string  `json:"specification"`                               // 规格
	Quantity      float64 `json:"quantity"`                                    // 数量
	UnitPrice     float64 `json:"unit_price"`                                  // 单价
	SupplierID    uint    `json:"supplier_id"`                                 // 供应商(外键)
	PaidMoney     float64 `json:"paid_money"`                                  // 已付金额
	PayMethod     string  `json:"pay_method"`                                  // 付款方式
	Note          string  `json:"note"`                                        // 备注
	Location      string  `json:"location"`                                    // 存放地点
}

// GenRestockOrder 根据进货信息生成进货订单
func (p *RestockParams) GenRestockOrder() *RestockOrder {
	return &RestockOrder{
		Date:          p.Date,
		ModelNum:      p.ModelNum,
		Specification: p.Specification,
		Quantity:      p.Quantity,
		UnitPrice:     p.UnitPrice,
		SumMoney:      p.Quantity * p.UnitPrice,
		SupplierID:    p.SupplierID,
		PayMethod:     p.PayMethod,
		Note:          p.Note,
	}
}
