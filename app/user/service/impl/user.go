package impl

import (
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
)

type UserImpl struct{}

func (UserImpl) Login(username string, password string) error {
	if username == global.Conf.Username && password == global.Conf.Password {
		return nil
	} else {
		return errors.WithMessagef(global.ErrUsernamePasswd, "账号密码验证错误")
	}
}
