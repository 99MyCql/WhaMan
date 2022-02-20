package do

import (
	restockDO "WhaMan/app/restock/do"

	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name          string                    `gorm:"unique;type:varchar(100);"` // 供应商名
	Contacts      string                    `gorm:"type:varchar(100);"`        // 联系人
	Phone         string                    `gorm:"type:varchar(100);"`        // 联系电话
	Turnover      float64                   // 交易额
	UnpaidMoney   float64                   // 未付款
	Note          string                    `gorm:"type:text"` // 备注
	RestockOrders []*restockDO.RestockOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
