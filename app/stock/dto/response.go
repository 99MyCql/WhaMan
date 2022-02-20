package dto

import (
	"time"

	restockDO "WhaMan/app/restock/do"
	sellDO "WhaMan/app/sell/do"
	"WhaMan/pkg/datetime"
)

// ComRsp Get List 等接口的响应数据
type ComRsp struct {
	ID              uint                    `json:"id"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	ModelNum        string                  `json:"model_num"`
	Specification   string                  `json:"specification"`
	RestockQuantity float64                 `json:"restock_quantity"`
	SellQuantity    float64                 `json:"sell_quantity"`
	CurQuantity     float64                 `json:"cur_quantity"`
	UnitPrice       float64                 `json:"unit_price"`
	SumMoney        float64                 `json:"sum_money"`
	Location        string                  `json:"location"`
	Note            string                  `json:"note"`
	RestockOrder    *restockDO.RestockOrder `json:"restock_order"`
	SellOrders      []*sellDO.SellOrder     `json:"sell_orders"`
	RestockDate     datetime.MyDatetime     `json:"restock_date"`
	SupplierName    string                  `json:"supplier_name"`
}
