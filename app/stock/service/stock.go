package service

import (
	"WhaMan/app/stock/model"

	"gorm.io/gorm"
)

type Stock interface {
	Find(id uint) (map[string]interface{}, error)
	List(option *model.ListOption) ([]map[string]interface{}, error)
	Update(id uint, p *model.UpdateParams) error
	UpdateSellOrders(tx *gorm.DB, id uint, unitPrice float64) error
}
