package service

import "WhaMan/app/sell/model"

// Sell 出货模块
type Sell interface {
	Sell(*model.SellParams) error
	Find(id uint) (map[string]interface{}, error)
	List(option *model.ListOption) ([]*model.SellOrder, error)
	Update(id uint, p *model.SellParams) error
	Delete(id uint) error
	// TODO: 分析每个月的交易情况
}
