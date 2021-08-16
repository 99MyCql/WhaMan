package model

import (
	sellModel "WhaMan/app/sell/model"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model  `binding:"-"`
	Name        string                 `gorm:"unique;type:varchar(100);" binding:"required,excludes= "` // 客户名
	Contacts    string                 `gorm:"type:varchar(100);" binding:"excludes= "`                 // 联系人
	Phone       string                 `gorm:"type:varchar(100);" binding:"excludes= ,len=11"`          // 联系电话
	Turnover    float64                `binding:"-"`                                                    // 交易额
	UnpaidMoney float64                `binding:"-"`                                                    // 未付款
	Note        string                 `gorm:"type:text" binding:"excludes= "`                          // 备注
	SellOrders  []*sellModel.SellOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" binding:"-"`
}
