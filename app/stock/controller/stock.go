package controller

import (
	"errors"
	"net/http"
	"strconv"

	"WhaMan/app/stock/model"
	"WhaMan/app/stock/service"
	"WhaMan/app/stock/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/gin-gonic/gin"
)

var stockService service.Stock = new(impl.StockImpl)

// @Summary Get
// @Tags Stock
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /stock/get/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)

	stock, err := stockService.Find(uint(id))
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.GetFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(stock))
}

// @Summary List
// @Tags Stock
// @Accept json
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /stock/list [post]
func List(c *gin.Context) {
	stocks, err := stockService.List()
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ListFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(stocks))
}

// @Summary Update
// @Tags Stock
// @Accept json
// @Param id path uint true "id"
// @Param data body model.UpdateParams true "更新的库存信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /stock/update/{id} [post]
func Update(c *gin.Context) {
	// 解析请求数据
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

	if err := stockService.Update(uint(id), req); err != nil {
		if errors.Is(err, global.ErrNameExist) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.UpdateFailed, "客户名称已存在"))
			return
		}
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.UpdateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary Delete
// @Tags Stock
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /stock/delete/{id} [get]
// func Delete(c *gin.Context) {
// 	// 解析请求数据
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		global.Log.Error(err)
// 		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
// 		return
// 	}
//
// 	if err := stockService.Delete(uint(id)); err != nil {
// 		global.Log.Errorf("%+v", err)
// 		if errors.Is(err, global.ErrCannotDelete) {
// 			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.DeleteFailed, "该库存存在出货订单，不能删除"))
// 			return
// 		}
// 		c.JSON(http.StatusOK, rsp.Err(rsp.DeleteFailed))
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, rsp.Suc())
// }
