package service

import (
	"WhaMan/app/stock/model"

	"gorm.io/gorm"
)

type Stock interface {
	Find(id uint) (*model.Stock, error)
	List() ([]*model.Stock, error)
	Update(id uint, p *model.UpdateParams) error
	UpdateSellOrders(tx *gorm.DB, id uint, unitPrice float64) error
	// Delete(id uint) error
}
