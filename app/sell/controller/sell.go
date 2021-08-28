package controller

import (
	"net/http"
	"strconv"

	"WhaMan/app/sell/model"
	"WhaMan/app/sell/service"
	"WhaMan/app/sell/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var sellService service.Sell = new(impl.SellImpl)

// @Summary Sell
// @Tags Sell
// @Accept json
// @Param data body model.SellParams true "出货信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /sell/sell [post]
func Sell(c *gin.Context) {
	var req *model.SellParams
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	err := sellService.Sell(req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.SellFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary GetSellOrder
// @Tags Sell
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /sell/getSellOrder/{id} [get]
func GetSellOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)

	sellOrder, err := sellService.Find(uint(id))
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.GetFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(sellOrder))
}

// @Summary ListSellOrders
// @Tags Sell
// @Accept json
// @Param data body model.ListOption true "选项"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /sell/listSellOrders [post]
func ListSellOrders(c *gin.Context) {
	var req *model.ListOption
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	sellOrders, err := sellService.List(req)
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ListFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(sellOrders))
}

// @Summary UpdateSellOrder
// @Tags Sell
// @Accept json
// @Param id path uint true "id"
// @Param data body model.SellParams true "出货订单信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /sell/updateSellOrder/{id} [post]
func UpdateRestockOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)
	var req *model.SellParams
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	if err := sellService.Update(uint(id), req); err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.UpdateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary DeleteSellOrder
// @Tags Sell
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /sell/deleteSellOrder/{id} [get]
func DeleteSellOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}

	if err := sellService.Delete(uint(id)); err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.DeleteFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
