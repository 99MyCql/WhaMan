package do

import (
	"WhaMan/pkg/datetime"

	"gorm.io/gorm"
)

type SellOrder struct {
	gorm.Model
	Date            datetime.MyDatetime `gorm:"not null;type:datetime"`      // 日期
	CustomerOrderID string              `gorm:"type:varchar(100);"`          // 客户订单号
	CustomerBatchID string              `gorm:"type:varchar(100);"`          // 客户批号
	DeliverOrderID  string              `gorm:"type:varchar(100);"`          // 送货单号
	ModelNum        string              `gorm:"not null;type:varchar(100);"` // 型号
	Specification   string              `gorm:"type:varchar(100);"`          // 规格
	Quantity        float64             `gorm:"not null"`                    // 数量
	UnitPrice       float64             `gorm:"not null"`                    // 出货单价
	PaidMoney       float64             `gorm:"not null"`                    // 已付金额（客户已经支付的金额）
	PayDate         datetime.MyDatetime `gorm:"type:datetime;default:null"`  // 付款日期
	PayMethod       string              `gorm:"type:varchar(100);"`          // 付款方式
	FreightCost     float64             // 运费
	Kickback        float64             // 回扣
	Tax             float64             // 税金
	OtherCost       float64             // 杂费
	Note            string              `gorm:"type:text"` // 备注

	// 可通过 RestockOrderID 推出，此处冗余是为了节省查询时间；另外，当 RestockOrderID 为null时，需指定进货单价
	RestockUnitPrice float64 `gorm:"not null"`     // 进货单价
	RestockOrderID   *uint   `gorm:"default:null"` // 进货订单编号(外键)，可为null
	CustomerID       uint    // 客户编号(外键)
}

// GetTurnover 获取交易额
func (s *SellOrder) GetTurnover() float64 {
	return s.UnitPrice * s.Quantity
}

// GetProfit 计算利润
func (s *SellOrder) GetProfit() float64 {
	return (s.UnitPrice-s.RestockUnitPrice)*s.Quantity - s.FreightCost - s.Kickback - s.Tax - s.OtherCost
}
