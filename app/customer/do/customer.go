package do

import (
	sellDO "WhaMan/app/sell/do"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string              `gorm:"unique;type:varchar(100);"` // 客户名
	Contacts   string              `gorm:"type:varchar(100);"`        // 联系人
	Phone      string              `gorm:"type:varchar(100);"`        // 联系电话
	Note       string              `gorm:"type:text"`                 // 备注
	SellOrders []*sellDO.SellOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
