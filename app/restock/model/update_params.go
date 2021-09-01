package model

import "WhaMan/pkg/global/models"

// UpdateParams 更新接口参数
type UpdateParams struct {
	Date          models.MyDatetime `json:"date" binding:"required"`      // 日期(字符串形式)
	ModelNum      string            `json:"model_num" binding:"required"` // 型号
	Specification string            `json:"specification"`                // 规格
	Quantity      float64           `json:"quantity"`                     // 数量
	UnitPrice     float64           `json:"unit_price"`                   // 单价
	PaidMoney     float64           `json:"paid_money"`                   // 已付金额
	PayMethod     string            `json:"pay_method"`                   // 付款方式
	Note          string            `json:"note"`                         // 备注
	SupplierID    uint              `json:"supplier_id"`                  // 供应商(外键)
}

func (p *UpdateParams) GenRestockOrder() *RestockOrder {
	return &RestockOrder{
		Date:          p.Date,
		ModelNum:      p.ModelNum,
		Specification: p.Specification,
		Quantity:      p.Quantity,
		UnitPrice:     p.UnitPrice,
		SumMoney:      p.Quantity * p.UnitPrice,
		PayMethod:     p.PayMethod,
		Note:          p.Note,
		SupplierID:    p.SupplierID,
	}
}
