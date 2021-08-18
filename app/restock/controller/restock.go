package controller

import (
	"net/http"
	"strconv"

	"WhaMan/app/restock/model"
	"WhaMan/app/restock/service"
	"WhaMan/app/restock/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

var restockService service.Restock = new(impl.RestockImpl)

// @Summary Restock
// @Tags Restock
// @Accept json
// @Param data body model.RestockParams true "进货信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /restock/restock [post]
func Restock(c *gin.Context) {
	// 解析请求数据
	var req *model.RestockParams
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}
	global.Log.Debugf("%+v", req)

	err := restockService.Restock(req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary GetRestockOrder
// @Tags Restock
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /restock/getRestockOrder/{id} [get]
func GetRestockOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)

	restockOrder, err := restockService.Find(uint(id))
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.GetFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(restockOrder))
}

// @Summary ListRestockOrders
// @Tags Restock
// @Accept json
// @Param data body model.ListOption true "选项"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /restock/listRestockOrders [post]
func ListRestockOrders(c *gin.Context) {
	var req *model.ListOption
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ParamError))
		return
	}
	global.Log.Debugf("%+v", req)

	restocks, err := restockService.List(req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ListFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(restocks))
}

// @Summary UpdateRestockOrder
// @Tags Restock
// @Accept json
// @Param id path uint true "id"
// @Param data body model.UpdateParams true "进货订单信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /restock/updateRestockOrder/{id} [post]
func UpdateRestockOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)
	var req *model.UpdateParams
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	if err := restockService.Update(uint(id), req); err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.UpdateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary DeleteRestockOrder
// @Tags Restock
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /restock/deleteRestockOrder/{id} [get]
func DeleteRestockOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}

	if err := restockService.Delete(uint(id)); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrCannotDelete) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.DeleteFailed, "该进货订单对应的库存已存在出货订单，不能删除"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.DeleteFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
