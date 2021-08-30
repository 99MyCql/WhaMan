package service

import (
	"WhaMan/app/stock/model"
)

type Stock interface {
	Find(id uint) (map[string]interface{}, error)
	List() ([]map[string]interface{}, error)
	Update(id uint, p *model.UpdateParams) error
}
