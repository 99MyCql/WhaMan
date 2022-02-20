package error

import (
	"fmt"
)

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

func (e *Error) AddMsg(msg string) *Error {
	e.Msg += " - " + msg
	return e
}
