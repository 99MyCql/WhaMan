package dto

import "WhaMan/app/customer/do"

// ComReq 创建、更新的请求参数
type ComReq struct {
	Name     string `binding:"required"` // 客户名
	Phone    string // 联系电话
	Contacts string // 联系人
	Note     string // 备注
}

func (r *ComReq) Convert2Customer() *do.Customer {
	return &do.Customer{
		Name:     r.Name,
		Contacts: r.Contacts,
		Phone:    r.Phone,
		Note:     r.Note,
	}
}
