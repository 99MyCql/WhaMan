package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	Unpaid   = "未结账"
	Paid     = "已结账"
	Returned = "已退货"
)

type SellOrder struct {
	ID              string    `gorm:"primarykey;type:varchar(100);"` // 订单号
	Date            time.Time `gorm:"not null"`                      // 日期
	State           string    `gorm:"not null;type:varchar(100);"`   // 状态
	CustomerOrderID string    `gorm:"type:varchar(100);"`            // 客户订单号
	CustomerBatchID string    `gorm:"type:varchar(100);"`            // 客户批号
	DeliverOrderID  string    `gorm:"type:varchar(100);"`            // 送货单号
	StockID         string    `gorm:"type:varchar(100);"`            // 库存号(外键)
	Specification   string    `gorm:"not null;type:varchar(100);"`   // 规格
	Quantity        float64   `gorm:"not null"`                      // 数量
	Unit            string    `gorm:"default:KG;type:varchar(100);"` // 单位
	UnitPrice       float64   `gorm:"not null"`                      // 单价
	SumMoney        float64   `gorm:"not null"`                      // 金额
	CustomerName    string    `gorm:"type:varchar(100);"`            // 客户(外键)
	PaidMoney       float64   `gorm:"not null"`                      // 已付金额
	PayMethod       string    `gorm:"type:varchar(100);"`            // 付款方式
	FreightCost     float64   // 运费
	Kickback        float64   // 回扣
	Tax             float64   // 税金
	OtherCost       float64   // 杂费
	Profit          float64   `gorm:"not null"`  // 利润
	Note            string    `gorm:"type:text"` // 备注
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
