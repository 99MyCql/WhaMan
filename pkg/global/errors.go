package global

import "github.com/pkg/errors"

var (
	ErrNameExist    = errors.New("名称已存在")
	ErrCannotDelete = errors.New("存在关联数据，不能删除")
)
