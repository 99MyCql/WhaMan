package model

import (
	"time"

	restockModel "WhaMan/app/restock/model"

	"gorm.io/gorm"
)

type Supplier struct {
	Name          string  `gorm:"primarykey;type:varchar(100);"` // 供应商名
	Contacts      string  `gorm:"type:varchar(100);"`            // 联系人
	Phone         string  `gorm:"type:varchar(100);"`            // 联系电话
	Turnover      float64 // 交易额
	Note          string  `gorm:"type:text"` // 备注
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt               `gorm:"index"`
	RestockOrders []*restockModel.RestockOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
