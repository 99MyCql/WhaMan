package model

import (
	sellModel "WhaMan/app/sell/model"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string                 `gorm:"unique;type:varchar(100);"` // 客户名
	Contacts    string                 `gorm:"type:varchar(100);"`        // 联系人
	Phone       string                 `gorm:"type:varchar(100);"`        // 联系电话
	Turnover    float64                // 交易额
	UnpaidMoney float64                // 未付款
	Note        string                 `gorm:"type:text"` // 备注
	SellOrders  []*sellModel.SellOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
