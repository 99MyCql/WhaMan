package model

import (
	restockModel "WhaMan/app/restock/model"
	sellModel "WhaMan/app/sell/model"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	ModelNum        string                     `json:"modelNum" gorm:"not null;type:varchar(100);"` // 型号
	Specification   string                     `json:"specification" gorm:"type:varchar(100);"`     // 规格
	RestockQuantity float64                    `json:"restockQuantity" gorm:"not null"`             // 进货数量
	SellQuantity    float64                    `json:"sellQuantity" gorm:"not null"`                // 出货数量
	CurQuantity     float64                    `json:"curQuantity" gorm:"not null"`                 // 库存数量
	UnitPrice       float64                    `json:"unitPrice" gorm:"not null"`                   // 单价
	SumMoney        float64                    `json:"sumMoney" gorm:"not null"`                    // 金额
	Location        string                     `json:"location" gorm:"type:varchar(100);"`          // 存放地点
	Note            string                     `json:"note" gorm:"type:text"`                       // 备注
	RestockOrder    *restockModel.RestockOrder `json:"restockOrder" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SellOrders      []*sellModel.SellOrder     `json:"sellOrders" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
