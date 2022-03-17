package dto

import (
	"encoding/json"

	"WhaMan/app/customer/do"
)

// ComReq 创建、更新的请求参数
type ComReq struct {
	Name     string `json:"name" binding:"required"` // 客户名
	Phone    string `json:"phone"`                   // 联系电话
	Contacts string `json:"contacts"`                // 联系人
	Note     string `json:"note"`                    // 备注
}

func (r *ComReq) Convert2Customer() *do.Customer {
	return &do.Customer{
		Name:     r.Name,
		Contacts: r.Contacts,
		Phone:    r.Phone,
		Note:     r.Note,
	}
}

type Date struct {
	StartDate string `json:"start_date" binding:"datetime=2006-01-02"`
	EndDate   string `json:"end_date" binding:"datetime=2006-01-02"`
}

type SellOrdersWhere struct {
	Date *Date `json:"date"`
}

type Where struct {
}

type ListReq struct {
	Where           *Where           `json:"where"`
	OrderBy         string           `json:"order_by"`
	SellOrdersWhere *SellOrdersWhere `json:"sell_orders_where"`
}

func (r *ListReq) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
