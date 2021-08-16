package global

import "github.com/pkg/errors"

var (
	ErrNameExist = errors.New("名称已存在")
)
