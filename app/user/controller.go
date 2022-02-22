package user

import (
	"WhaMan/app/user/dto"
	myErr "WhaMan/pkg/error"
	"WhaMan/pkg/log"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var service = new(Service)

// @Summary Login
// @Tags User
// @Accept json
// @Param data body dto.LoginReq true "用户信息"
// @Success 200 {object} rsp.ResponseExample "code=200"
// @Failure 400 {object} rsp.ResponseExample "code=4xxxxx"
// @Failure 500 {object} rsp.ResponseExample "code=5xxxxx"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var req *dto.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		log.Logger.Error(err)
		c.JSON(rsp.Err(myErr.ParamErr))
		return
	}

	c.JSON(rsp.New(service.Login(req, c)))
}
