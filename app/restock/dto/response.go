package dto

import (
	"time"

	"WhaMan/pkg/datetime"
)

type ComRsp struct {
	ID            uint                `json:"id"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	Date          datetime.MyDatetime `json:"date"` // 日期(2000-01-01)
	ModelNum      string              `json:"model_num"`
	Specification string              `json:"specification"`
	Quantity      float64             `json:"quantity"`
	UnitPrice     float64             `json:"unit_price"`
	SumMoney      float64             `json:"sum_money"`
	PaidMoney     float64             `json:"paid_money"`
	PayMethod     string              `json:"pay_method"`
	Location      string              `json:"location"`
	Note          string              `json:"note"`
	StockID       uint                `json:"stock_id"`    // 库存编号(外键)
	SupplierID    uint                `json:"supplier_id"` // 供应商编号(外键)
	SupplierName  string              `json:"supplier_name"`
}
