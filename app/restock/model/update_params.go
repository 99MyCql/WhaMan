package model

// UpdateParams 更新接口参数
type UpdateParams struct {
	Date          string  `binding:"datetime=2006-01-02"` // 日期(字符串形式)
	ModelNum      string  `binding:"required"`            // 型号
	Specification string  `binding:"required"`            // 规格
	Quantity      float64 `binding:"required,ne=0"`       // 数量
	UnitPrice     float64 `binding:"required,ne=0"`       // 单价
	PaidMoney     float64 // 已付金额
	PayMethod     string  // 付款方式
	Note          string  // 备注
	SupplierID    uint    // 供应商(外键)
}

func (p *UpdateParams) GenRestockOrder() *RestockOrder {
	return &RestockOrder{
		Date:          p.Date,
		ModelNum:      p.ModelNum,
		Specification: p.Specification,
		Quantity:      p.Quantity,
		UnitPrice:     p.UnitPrice,
		SumMoney:      p.Quantity * p.UnitPrice,
		PaidMoney:     p.PaidMoney,
		PayMethod:     p.PayMethod,
		Note:          p.Note,
		SupplierID:    p.SupplierID,
	}
}
