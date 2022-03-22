package dto

import (
	"time"

	SellDTO "WhaMan/app/sell/dto"
)

type GetRsp struct {
	ID         uint              `json:"id"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Name       string            `json:"name"`
	Contacts   string            `json:"contacts"`
	Phone      string            `json:"phone"`
	Note       string            `json:"note"`
	SellOrders []*SellDTO.ComRsp `json:"sell_orders" gorm:"-"`
}

type ListRsp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Contacts  string    `json:"contacts"`
	Phone     string    `json:"phone"`
	Turnover  float64   `json:"turnover"`
	PaidMoney float64   `json:"paid_money"`
	Profit    float64   `json:"profit"`
	Note      string    `json:"note"`
}
