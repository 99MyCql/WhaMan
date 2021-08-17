package model

// UpdateParams 更新接口的输入参数
type UpdateParams struct {
	Location string // 存放地点
	Note     string // 备注
}

func (p *UpdateParams) GenStock() *Stock {
	return &Stock{
		Location: p.Location,
		Note:     p.Note,
	}
}
