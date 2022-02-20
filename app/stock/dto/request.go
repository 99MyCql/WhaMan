package dto

import (
	"encoding/json"

	"WhaMan/app/stock/do"
)

// UpdateReq 更新接口的请求参数
type UpdateReq struct {
	Location string // 存放地点
	Note     string // 备注
}

func (p *UpdateReq) Convert2Stock() *do.Stock {
	return &do.Stock{
		Location: p.Location,
		Note:     p.Note,
	}
}

// ListReq List 接口的请求参数
type ListReq struct {
	Where *struct {
		CurQuantity *struct {
			Start *float64
			End   *float64
		}
	}
	OrderBy string
}

func (o *ListReq) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
