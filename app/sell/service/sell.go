package service

import "WhaMan/app/sell/model"

// Sell 出货模块
type Sell interface {
	Sell(*model.SellParams) error
	Find(id uint) (*model.SellOrder, error)
	List(option *model.ListOption) ([]*model.SellOrder, error)
	Update(id uint, p *model.SellParams) error
	Delete(id uint) error
}
