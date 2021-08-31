package model

import "encoding/json"

type ListOption struct {
	Where *struct {
		Date *struct {
			StartDate string `binding:"datetime=2006-01-02"`
			EndDate   string `binding:"datetime=2006-01-02"`
		}
		CustomerID uint
		StockID    uint
	}
	OrderBy string
}

func (o *ListOption) String() string {
	out, err := json.Marshal(o)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
