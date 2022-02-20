package dto

import (
	"time"
)

type ComRsp struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Contacts    string    `json:"contacts"`
	Phone       string    `json:"phone"`
	Turnover    float64   `json:"turnover"`
	UnpaidMoney float64   `json:"unpaid_money"`
	Note        string    `json:"note"`
}
