package controller

import (
	"errors"
	"net/http"
	"strconv"

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
// @Param data body model.Params true "客户信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /customer/create [post]
func Create(c *gin.Context) {
	// 解析请求数据
	var req *model.Params
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	if err := customerService.Create(req); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrNameExist) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.CreateFailed, "客户名称已存在"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary Get
// @Tags Customer
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /customer/get/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)

	customer, err := customerService.Find(uint(id))
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.GetFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(customer))
}

// @Summary List
// @Tags Customer
// @Accept json
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /customer/list [post]
func List(c *gin.Context) {
	customers, err := customerService.List()
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ListFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(customers))
}

// @Summary Update
// @Tags Customer
// @Accept json
// @Param id path uint true "id"
// @Param data body model.Params true "客户信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /customer/update/{id} [post]
func Update(c *gin.Context) {
	// 解析请求数据
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)
	var req *model.Params
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	if err := customerService.Update(uint(id), req); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrNameExist) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.UpdateFailed, "客户名称已存在"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.UpdateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary Delete
// @Tags Customer
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /customer/delete/{id} [get]
func Delete(c *gin.Context) {
	// 解析请求数据
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}

	if err := customerService.Delete(uint(id)); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrCannotDelete) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.DeleteFailed, "与该客户存在交易订单，不能删除"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.DeleteFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
