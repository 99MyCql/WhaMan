package model

// RestockParams 进货接口参数
type RestockParams struct {
	Date          string  `binding:"required,datetime=2006-01-02"` // 日期(字符串形式)
	ModelNum      string  `binding:"required"`                     // 型号
	Specification string  `binding:"required"`                     // 规格
	Quantity      float64 `binding:"required,ne=0"`                // 数量
	UnitPrice     float64 `binding:"required,ne=0"`                // 单价
	SumMoney      float64 `binding:"required,ne=0"`                // 金额
	SupplierID    uint    `binding:"required,ne=0"`                // 供应商(外键)
	PaidMoney     float64 `binding:"required"`                     // 已付金额
	PayMethod     string  // 付款方式
	Note          string  // 备注
	Location      string  // 存放地点
}

// GenRestockOrder 根据进货信息生成进货订单
func (p *RestockParams) GenRestockOrder() *RestockOrder {
	return &RestockOrder{
		Date:          p.Date,
		ModelNum:      p.ModelNum,
		Specification: p.Specification,
		Quantity:      p.Quantity,
		UnitPrice:     p.UnitPrice,
		SumMoney:      p.SumMoney,
		SupplierID:    p.SupplierID,
		PaidMoney:     p.PaidMoney,
		PayMethod:     p.PayMethod,
		Note:          p.Note,
	}
}
