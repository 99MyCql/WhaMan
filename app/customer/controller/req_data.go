package controller

import (
	"WhaMan/app/customer/model"
)

// reqData 用于创建、更新等接口解析请求数据
type reqData struct {
	Name     string `binding:"required,excludes= "` // 客户名
	Contacts string `binding:"excludes= "`          // 联系人
	Phone    string `binding:"excludes= "`          // 联系电话
	Note     string `binding:"excludes= "`          // 备注
}

func (r reqData) genCustomer() *model.Customer {
	return &model.Customer{
		Name:     r.Name,
		Contacts: r.Contacts,
		Phone:    r.Phone,
		Note:     r.Note,
	}
}
