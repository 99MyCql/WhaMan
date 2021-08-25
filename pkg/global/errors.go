package global

import "github.com/pkg/errors"

var (
	ErrNameExist      = errors.New("名称已存在")
	ErrCannotDelete   = errors.New("存在关联数据，不能删除")
	ErrUsernamePasswd = errors.New("账号或密码错误")
)
