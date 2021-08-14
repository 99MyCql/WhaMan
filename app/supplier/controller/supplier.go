package controller

import (
	"net/http"

	"WhaMan/app/supplier/model"
	"WhaMan/app/supplier/service"
	"WhaMan/app/supplier/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var supplierService service.Supplier = new(impl.Supplier)

// @Summary Create
// @Tags Supplier
// @Accept json
// @Param supplierInfo body model.SupplierInfo true "supplierInfo"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /supplier/create [post]
func Create(c *gin.Context) {
	// 定义请求数据结构
	var req *model.SupplierInfo
	// 解析请求数据
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}
	global.Log.Debug(req)

	err := supplierService.Create(req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
