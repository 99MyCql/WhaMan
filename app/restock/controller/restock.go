package controller

import (
	"net/http"
	"time"

	"WhaMan/app/restock/model"
	"WhaMan/app/restock/service"
	"WhaMan/app/restock/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var restockService service.Restock = new(impl.Restock)

// Restock 进货：新增进货订单，并新增库存(一个库存对应一个进货订单)
// @Summary Restock
// @Tags Restock
// @Accept json
// @Param RestockInfo body model.RestockInfo true "RestockInfo"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /restock/restock [post]
func Restock(c *gin.Context) {
	// 定义请求数据结构
	var req model.RestockInfo
	// 解析请求数据
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}
	global.Log.Debugf("%+v", req)

	// 将string格式的时间转换为time.Time格式的时间
	var err error
	req.Date, err = time.Parse("2006/01/02", req.DateStr)
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}

	err = restockService.Restock(&req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
