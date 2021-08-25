package controller

import (
	"errors"
	"net/http"

	"WhaMan/app/user/service"
	"WhaMan/app/user/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

var userService service.User = new(impl.UserImpl)

type loginReq struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

// @Summary Login
// @Tags User
// @Accept json
// @Param data body loginReq true "用户信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var req *loginReq
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	err := userService.Login(req.Username, req.Password)
	if err != nil {
		global.Log.Error(err)
		if errors.Is(err, global.ErrUsernamePasswd) {
			c.JSON(http.StatusOK, rsp.Err(rsp.UsernamePasswdError))
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.LoginFailed))
		return
	}

	session := sessions.Default(c)
	session.Set("isLogin", true)
	if err := session.Save(); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.LoginFailed))
		return
	}
	c.JSON(http.StatusOK, rsp.Suc())
}
