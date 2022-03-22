package error

import (
	"fmt"
)

type Error struct {
	Code   int
	Msg    string
	Detail string
}

func (e *Error) Error() string {
	if e.Detail == "" {
		return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
	}
	return fmt.Sprintf("[%d] %s:%s", e.Code, e.Msg, e.Detail)
}

func (e *Error) SetDetail(detail string) *Error {
	e.Detail = detail
	return e
}
