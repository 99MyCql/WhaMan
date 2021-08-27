package model

import (
	restockModel "WhaMan/app/restock/model"

	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name          string                       `json:"name" gorm:"unique;type:varchar(100);"` // 供应商名
	Contacts      string                       `json:"contacts" gorm:"type:varchar(100);"`    // 联系人
	Phone         string                       `json:"phone" gorm:"type:varchar(100);"`       // 联系电话
	Turnover      float64                      `json:"turnover"`                              // 交易额
	Note          string                       `json:"note" gorm:"type:text"`                 // 备注
	RestockOrders []*restockModel.RestockOrder `json:"restock_orders" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
