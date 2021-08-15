package model

import (
	"time"

	"gorm.io/gorm"
)

type RestockOrder struct {
	gorm.Model
	Date          time.Time `gorm:"not null"`                      // 日期
	ModelNum      string    `gorm:"not null;type:varchar(100);"`   // 型号
	Specification string    `gorm:"not null;type:varchar(100);"`   // 规格
	Quantity      float64   `gorm:"not null;type:varchar(100);"`   // 数量
	Unit          string    `gorm:"default:KG;type:varchar(100);"` // 单位
	UnitPrice     float64   `gorm:"not null"`                      // 单价
	SumMoney      float64   `gorm:"not null"`                      // 金额
	PaidMoney     float64   `gorm:"not null"`                      // 已付金额
	PayMethod     string    `gorm:"not null;type:varchar(100);"`   // 付款方式
	Note          string    `gorm:"type:text"`                     // 备注
	StockID       uint      // 库存编号(外键)
	SupplierID    uint      // 供应商编号(外键)
}
