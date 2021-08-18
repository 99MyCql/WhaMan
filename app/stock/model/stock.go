package model

import (
	restockModel "WhaMan/app/restock/model"
	sellModel "WhaMan/app/sell/model"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	ModelNum        string                     `gorm:"not null;type:varchar(100);"` // 型号
	Specification   string                     `gorm:"not null;type:varchar(100);"` // 规格
	RestockQuantity float64                    `gorm:"not null"`                    // 进货数量
	SellQuantity    float64                    `gorm:"not null"`                    // 出货数量
	CurQuantity     float64                    `gorm:"not null"`                    // 库存数量
	UnitPrice       float64                    `gorm:"not null"`                    // 单价
	SumMoney        float64                    `gorm:"not null"`                    // 金额
	Location        string                     `gorm:"type:varchar(100);"`          // 存放地点
	Note            string                     `gorm:"type:text"`                   // 备注
	RestockOrder    *restockModel.RestockOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SellOrders      []*sellModel.SellOrder     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Unit            string                     `gorm:"default:KG;type:varchar(100);"` // 单位
}
