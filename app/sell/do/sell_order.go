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
	// TODO: paidMoney已付款改为receivedMoney已收款
	gorm.Model
	Date             datetime.MyDatetime `json:"date" gorm:"not null;type:datetime"`               // 日期
	CustomerOrderID  string              `json:"customerOrderID" gorm:"type:varchar(100);"`        // 客户订单号
	CustomerBatchID  string              `json:"customerBatchID" gorm:"type:varchar(100);"`        // 客户批号
	DeliverOrderID   string              `json:"deliverOrderID" gorm:"type:varchar(100);"`         // 送货单号
	ModelNum         string              `json:"modelNum" gorm:"not null;type:varchar(100);"`      // 型号
	Specification    string              `json:"specification" gorm:"not null;type:varchar(100);"` // 规格
	Quantity         float64             `json:"quantity" gorm:"not null"`                         // 数量
	RestockUnitPrice float64             `json:"restockUnitPrice" gorm:"not null"`                 // 进货单价
	UnitPrice        float64             `json:"unitPrice" gorm:"not null"`                        // 出货单价
	SumMoney         float64             `json:"sumMoney" gorm:"not null"`                         // 金额
	PaidMoney        float64             `json:"paidMoney" gorm:"not null"`                        // 已付金额
	PayDate          datetime.MyDatetime `json:"payDate" gorm:"type:datetime;default:null"`        // 付款日期
	PayMethod        string              `json:"payMethod" gorm:"type:varchar(100);"`              // 付款方式
	FreightCost      float64             `json:"freightCost"`                                      // 运费
	Kickback         float64             `json:"kickback"`                                         // 回扣
	Tax              float64             `json:"tax"`                                              // 税金
	OtherCost        float64             `json:"otherCost"`                                        // 杂费
	Profit           float64             `json:"profit" gorm:"not null"`                           // 利润
	Note             string              `json:"note" gorm:"type:text"`                            // 备注
	StockID          uint                `json:"stockID" gorm:"default:null"`                      // 库存编号(外键)
	CustomerID       uint                `json:"customerID"`                                       // 客户编号(外键)
}

// CalProfit 计算利润
func (s *SellOrder) CalProfit() {
	s.Profit = s.Quantity*(s.UnitPrice-s.RestockUnitPrice) - s.FreightCost - s.Kickback - s.Tax - s.OtherCost
}
