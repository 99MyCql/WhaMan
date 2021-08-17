package model

// Params 创建、更新的参数
type Params struct {
	Name     string `binding:"required,excludes= "` // 客户名
	Contacts string `binding:"excludes= "`          // 联系人
	Phone    string `binding:"excludes= "`          // 联系电话
	Note     string `binding:"excludes= "`          // 备注
}

func (p *Params) GenCustomer() *Customer {
	return &Customer{
		Name:     p.Name,
		Contacts: p.Contacts,
		Phone:    p.Phone,
		Note:     p.Note,
	}
}
