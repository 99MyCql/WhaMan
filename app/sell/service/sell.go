package service

import "WhaMan/app/sell/model"

// Sell 出货模块
type Sell interface {
	Sell(*model.Params) error
	Find(id uint) (*model.SellOrder, error)
	List(option *model.ListOption) ([]*model.SellOrder, error)
	Update(id uint, p *model.Params) error
	Delete(id uint) error
}
