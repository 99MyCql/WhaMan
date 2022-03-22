package user

import (
	"WhaMan/app/user/dto"
	"WhaMan/config"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Service struct{}

func (Service) Login(req *dto.LoginReq, c *gin.Context) error {
	if req.Username != config.Conf.Username || req.Password != config.Conf.Password {
		return myErr.UsernamePasswdErr
	}
	session := sessions.Default(c)
	session.Set("isLogin", true)
	if err := session.Save(); err != nil {
		log.Logger.Error(err)
		return myErr.ServerErr
	}
	return nil
}
