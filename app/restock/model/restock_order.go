package model

import (
	"gorm.io/gorm"
)

type RestockOrder struct {
	gorm.Model
	Date          string  `json:"date" gorm:"not null;type:datetime"`          // 日期(2000-01-01)
	ModelNum      string  `json:"modelNum" gorm:"not null;type:varchar(100);"` // 型号
	Specification string  `json:"specification" gorm:"type:varchar(100);"`     // 规格
	Quantity      float64 `json:"quantity" gorm:"not null;type:varchar(100);"` // 数量
	UnitPrice     float64 `json:"unitPrice" gorm:"not null"`                   // 单价
	SumMoney      float64 `json:"sumMoney" gorm:"not null"`                    // 金额
	PayMethod     string  `json:"payMethod" gorm:"type:varchar(100);"`         // 付款方式
	Note          string  `json:"note" gorm:"type:text"`                       // 备注
	StockID       uint    `json:"stockID"`                                     // 库存编号(外键)
	SupplierID    uint    `json:"supplierID"`                                  // 供应商编号(外键)
}
