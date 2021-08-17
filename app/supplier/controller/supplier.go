package controller

import (
	"net/http"
	"strconv"

	"WhaMan/app/supplier/model"
	"WhaMan/app/supplier/service"
	"WhaMan/app/supplier/service/impl"
	"WhaMan/pkg/global"
	"WhaMan/pkg/rsp"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

var supplierService service.Supplier = new(impl.Supplier)

// @Summary Create
// @Tags Supplier
// @Accept json
// @Param data body model.Params true "供应商信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /supplier/create [post]
func Create(c *gin.Context) {
	// 解析请求数据
	var req *model.Params
	if err := c.ShouldBind(&req); err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debugf("%+v", req)

	if err := supplierService.Create(req); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrNameExist) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.UpdateFailed, "供应商名称已存在"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.CreateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary Get
// @Tags Supplier
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /supplier/get/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}
	global.Log.Debug(id)

	supplier, err := supplierService.Find(uint(id))
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.GetFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(supplier))
}

// @Summary List
// @Tags Supplier
// @Accept json
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /supplier/list [post]
func List(c *gin.Context) {
	suppliers, err := supplierService.List()
	if err != nil {
		global.Log.Errorf("%+v", err)
		c.JSON(http.StatusOK, rsp.Err(rsp.ListFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.SucWithData(suppliers))
}

// @Summary Update
// @Tags Supplier
// @Accept json
// @Param id path uint true "id"
// @Param data body model.Params true "供应商信息"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /supplier/update/{id} [post]
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

	if err := supplierService.Update(uint(id), req); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrNameExist) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.UpdateFailed, "供应商名称已存在"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.UpdateFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}

// @Summary Delete
// @Tags Supplier
// @Accept json
// @Param id path uint true "id"
// @Success 200 {string} json "{"code":0,"data":{},"msg":""}"
// @Failure 200 {string} json "{"code":非0,"data":{},"msg":""}"
// @Router /supplier/delete/{id} [get]
func Delete(c *gin.Context) {
	// 解析请求数据
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.ParamError, err.Error()))
		return
	}

	if err := supplierService.Delete(uint(id)); err != nil {
		global.Log.Errorf("%+v", err)
		if errors.Is(err, global.ErrCannotDelete) {
			c.JSON(http.StatusOK, rsp.ErrWithMsg(rsp.DeleteFailed, "与该供应商存在交易订单，不能删除"))
			return
		}
		c.JSON(http.StatusOK, rsp.Err(rsp.DeleteFailed))
		return
	}

	c.JSON(http.StatusOK, rsp.Suc())
}
