package dto

import (
	"time"

	"WhaMan/pkg/datetime"
)

type ComRsp struct {
	ID               uint                `json:"id"`
	CreatedAt        time.Time           `json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
	Date             datetime.MyDatetime `json:"date"`
	CustomerOrderID  string              `json:"customer_order_id"`
	CustomerBatchID  string              `json:"customer_batch_id"`
	DeliverOrderID   string              `json:"deliver_order_id"`
	ModelNum         string              `json:"model_num"`
	Specification    string              `json:"specification"`
	Quantity         float64             `json:"quantity"`
	UnitPrice        float64             `json:"unit_price"`
	PaidMoney        float64             `json:"paid_money"`
	PayDate          datetime.MyDatetime `json:"pay_date"`
	PayMethod        string              `json:"pay_method"`
	FreightCost      float64             `json:"freight_cost"`
	Kickback         float64             `json:"kickback"`
	Tax              float64             `json:"tax"`
	OtherCost        float64             `json:"other_cost"`
	Note             string              `json:"note"`
	RestockUnitPrice float64             `json:"restock_unit_price"`
	RestockOrderID   *uint               `json:"restock_order_id"`
	CustomerID       uint                `json:"customer_id"`
	CustomerName     string              `json:"customer_name"`
}
