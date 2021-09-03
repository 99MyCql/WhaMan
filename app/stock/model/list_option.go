package model

import "encoding/json"

type ListOption struct {
	Where *struct {
		CurQuantity *struct {
			Start *float64
			End   *float64
		}
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
