package controller

import (
	"WhaMan/app/supplier/model"
)

type reqData struct {
	Name     string `binding:"required,excludes= "` // 供应商名
	Contacts string `binding:"excludes= "`          // 联系人
	Phone    string `binding:"excludes= "`          // 联系电话
	Note     string `binding:"excludes= "`          // 备注
}

// genSupplier 根据 reqData 生成 Supplier
func (r *reqData) genSupplier() *model.Supplier {
	return &model.Supplier{
		Name:     r.Name,
		Contacts: r.Contacts,
		Phone:    r.Phone,
		Note:     r.Note,
	}
}
