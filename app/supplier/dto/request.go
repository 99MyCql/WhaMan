package dto

import (
	"encoding/json"

	"WhaMan/app/supplier/do"
)

// ComReq 用于 Create Update 接口的请求参数
type ComReq struct {
	Name     string `json:"name" binding:"required"` // 供应商名
	Phone    string `json:"phone"`                   // 联系电话
	Contacts string `json:"contacts"`                // 联系人
	Note     string `json:"note"`                    // 备注
}

func (r *ComReq) Convert2Supplier() *do.Supplier {
	return &do.Supplier{
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

type RestockOrdersWhere struct {
	Date *Date `json:"date"`
}

type Where struct {
}

type ListReq struct {
	Where              *Where              `json:"where"`
	OrderBy            string              `json:"order_by"`
	RestockOrdersWhere *RestockOrdersWhere `json:"restock_orders_where"`
}

func (r *ListReq) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
