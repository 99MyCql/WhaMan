package model

// Params 用于Create、Update接口的输入参数
type Params struct {
	Name     string `binding:"required,excludes= "` // 供应商名
	Contacts string `binding:"excludes= "`          // 联系人
	Phone    string `binding:"excludes= "`          // 联系电话
	Note     string `binding:"excludes= "`          // 备注
}

func (p *Params) GenSupplier() *Supplier {
	return &Supplier{
		Name:     p.Name,
		Contacts: p.Contacts,
		Phone:    p.Phone,
		Note:     p.Note,
	}
}