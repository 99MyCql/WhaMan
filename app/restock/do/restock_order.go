package do

import (
	sellDO "WhaMan/app/sell/do"
	"WhaMan/pkg/datetime"

	"gorm.io/gorm"
)

type RestockOrder struct {
	gorm.Model
	Date          datetime.MyDatetime `gorm:"not null;type:datetime"`      // 日期(2000-01-01)
	ModelNum      string              `gorm:"not null;type:varchar(100);"` // 型号
	Specification string              `gorm:"type:varchar(100);"`          // 规格
	Quantity      float64             `gorm:"not null"`                    // 数量
	UnitPrice     float64             `gorm:"not null"`                    // 单价
	PaidMoney     float64             `gorm:"not null"`                    // 已付金额
	PayMethod     string              `gorm:"type:varchar(100);"`          // 付款方式
	Location      string              `gorm:"type:varchar(100);"`          // 存放地点
	Note          string              `gorm:"type:text"`                   // 备注
	SupplierID    uint                // 供应商编号(外键)
	SellOrders    []*sellDO.SellOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// GetTurnover 获取订单交易额
func (r *RestockOrder) GetTurnover() float64 {
	return r.UnitPrice * r.Quantity
}
