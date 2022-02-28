package do

import (
	"WhaMan/pkg/datetime"

	"gorm.io/gorm"
)

// const (
// 	Unpaid   = "未结账"
// 	Paid     = "已结账"
// 	Returned = "已退货"
// )

type SellOrder struct {
	gorm.Model
	Date             datetime.MyDatetime `gorm:"not null;type:datetime"`      // 日期
	CustomerOrderID  string              `gorm:"type:varchar(100);"`          // 客户订单号
	CustomerBatchID  string              `gorm:"type:varchar(100);"`          // 客户批号
	DeliverOrderID   string              `gorm:"type:varchar(100);"`          // 送货单号
	ModelNum         string              `gorm:"not null;type:varchar(100);"` // 型号
	Specification    string              `gorm:"not null;type:varchar(100);"` // 规格
	Quantity         float64             `gorm:"not null"`                    // 数量
	RestockUnitPrice float64             `gorm:"not null"`                    // 进货单价
	UnitPrice        float64             `gorm:"not null"`                    // 出货单价
	SumMoney         float64             `gorm:"not null"`                    // 金额
	PaidMoney        float64             `gorm:"not null"`                    // 已付金额（客户已经支付的金额）
	PayDate          datetime.MyDatetime `gorm:"type:datetime;default:null"`  // 付款日期
	PayMethod        string              `gorm:"type:varchar(100);"`          // 付款方式
	FreightCost      float64             // 运费
	Kickback         float64             // 回扣
	Tax              float64             // 税金
	OtherCost        float64             // 杂费
	Profit           float64             `gorm:"not null"`     // 利润
	Note             string              `gorm:"type:text"`    // 备注
	StockID          *uint               `gorm:"default:null"` // 库存编号(外键)，可为null
	CustomerID       uint                // 客户编号(外键)
}

// CalProfit 计算利润
func (s *SellOrder) CalProfit() {
	s.Profit = s.Quantity*(s.UnitPrice-s.RestockUnitPrice) - s.FreightCost - s.Kickback - s.Tax - s.OtherCost
}
