package service

import "WhaMan/app/stock/model"

type Stock interface {
	Find(id uint) (*model.Stock, error)
	List() ([]*model.Stock, error)
	Update(id uint, p *model.UpdateParams) error
	// Delete(id uint) error
}
