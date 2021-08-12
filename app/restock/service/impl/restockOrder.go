package impl

import (
	restockModel "WhaMan/app/restock/model"
)

type RestockOrder struct{}

func (RestockOrder) Create(r *restockModel.RestockOrder) error {
	panic("implement me")
}

func (RestockOrder) Find(id string) (*restockModel.RestockOrder, error) {
	panic("implement me")
}

func (RestockOrder) List() ([]*restockModel.RestockOrder, error) {
	panic("implement me")
}

func (RestockOrder) Update(r *restockModel.RestockOrder) error {
	panic("implement me")
}

func (RestockOrder) Delete(id string) error {
	panic("implement me")
}
