package model

// Params 创建、更新的参数
type Params struct {
	Name     string `binding:"required"` // 客户名
	Phone    string // 联系电话
	Contacts string // 联系人
	Note     string // 备注
}

func (p *Params) GenCustomer() *Customer {
	return &Customer{
		Name:     p.Name,
		Contacts: p.Contacts,
		Phone:    p.Phone,
		Note:     p.Note,
	}
}
