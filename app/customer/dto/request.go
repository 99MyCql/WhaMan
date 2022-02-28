package dto

import "WhaMan/app/customer/do"

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
