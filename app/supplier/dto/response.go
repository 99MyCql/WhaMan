package dto

import (
	"time"

	restockDTO "WhaMan/app/restock/dto"
)

// GetRsp Get List 等接口响应数据
type GetRsp struct {
	ID            uint                 `json:"id"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Name          string               `json:"name"`
	Contacts      string               `json:"contacts"`
	Phone         string               `json:"phone"`
	Note          string               `json:"note"`
	RestockOrders []*restockDTO.GetRsp `json:"restock_orders" gorm:"-"`
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
	Note      string    `json:"note"`
}
