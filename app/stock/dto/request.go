package dto

import (
	"encoding/json"

	"WhaMan/app/stock/do"
)

// UpdateReq 更新接口的请求参数
type UpdateReq struct {
	Location string `json:"location"` // 存放地点
	Note     string `json:"note"`     // 备注
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
			Start *float64 `json:"start"`
			End   *float64 `json:"end"`
		} `json:"cur_quantity"`
	} `json:"where"`
	OrderBy string `json:"order_by"`
}

func (o *ListReq) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
