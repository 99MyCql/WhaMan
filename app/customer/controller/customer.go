package controller

import (
	"net/http"

	"WhaMan/app/customer/model"
	"WhaMan/app/customer/service"
	"WhaMan/app/customer/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var customerService service.Customer = new(impl.CustomerImpl)

// @Summary Create
// @Tags Customer
// @Accept json
// @Param customerInfo body model.CustomerInfo true "customerInfo"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /customer/create [post]
func Create(c *gin.Context) {
	// 定义请求数据结构
	var req *model.CustomerInfo
	// 解析请求数据
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}
	global.Log.Debug(req)

	err := customerService.Create(req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
