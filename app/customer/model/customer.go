package model

import (
	sellModel "WhaMan/app/sell/model"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string                 `json:"name" gorm:"unique;type:varchar(100);"` // 客户名
	Contacts    string                 `json:"contacts" gorm:"type:varchar(100);"`    // 联系人
	Phone       string                 `json:"phone" gorm:"type:varchar(100);"`       // 联系电话
	Turnover    float64                `json:"turnover"`                              // 交易额
	UnpaidMoney float64                `json:"unpaidMoney"`                           // 未付款
	Note        string                 `json:"note" gorm:"type:text"`                 // 备注
	SellOrders  []*sellModel.SellOrder `json:"sellOrders" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
