package do

import (
	restockDO "WhaMan/app/restock/do"
	sellDO "WhaMan/app/sell/do"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	ModelNum        string                  `gorm:"not null;type:varchar(100);"` // 型号
	Specification   string                  `gorm:"type:varchar(100);"`          // 规格
	RestockQuantity float64                 `gorm:"not null"`                    // 进货数量
	SellQuantity    float64                 `gorm:"not null"`                    // 出货数量
	CurQuantity     float64                 `gorm:"not null"`                    // 库存数量
	UnitPrice       float64                 `gorm:"not null"`                    // 单价
	SumMoney        float64                 `gorm:"not null"`                    // 金额
	Location        string                  `gorm:"type:varchar(100);"`          // 存放地点
	Note            string                  `gorm:"type:text"`                   // 备注
	RestockOrder    *restockDO.RestockOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SellOrders      []*sellDO.SellOrder     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
