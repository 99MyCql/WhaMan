package dto

import (
	"WhaMan/app/supplier/do"
)

// ComReq 用于 Create Update 接口的请求参数
type ComReq struct {
	Name     string `binding:"required"` // 供应商名
	Phone    string // 联系电话
	Contacts string // 联系人
	Note     string // 备注
}

func (r *ComReq) Convert2Supplier() *do.Supplier {
	return &do.Supplier{
		Name:     r.Name,
		Contacts: r.Contacts,
		Phone:    r.Phone,
		Note:     r.Note,
	}
}
