package main

import (
	customerController "WhaMan/app/customer/controller"
	restockController "WhaMan/app/restock/controller"
	sellController "WhaMan/app/sell/controller"
	stockController "WhaMan/app/stock/controller"
	supplierController "WhaMan/app/supplier/controller"
	_ "WhaMan/docs"
	"WhaMan/pkg/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/global"
	"WhaMan/pkg/log"
	"WhaMan/pkg/validators"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func init() {
	global.Log = log.New(log.DebugLevel)
	var err error
	if global.Conf, err = config.New("conf.yml"); err != nil {
		global.Log.Fatalf("%+v", err)
	}
	if global.DB, err = database.New(); err != nil {
		global.Log.Fatalf("%+v", err)
	}
}

func main() {
	r := gin.Default()

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("datetime", validators.DatetimeFormat); err != nil {
			global.Log.Fatal(err)
		}
	}

	// 注册 swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	restock := r.Group("/restock")
	{
		restock.POST("/restock", restockController.Restock)
		restock.GET("/getRestockOrder/:id", restockController.GetRestockOrder)
		restock.POST("/listRestockOrders", restockController.ListRestockOrders)
		restock.POST("/updateRestockOrder/:id", restockController.UpdateRestockOrder)
		restock.GET("/deleteRestockOrder/:id", restockController.DeleteRestockOrder)
	}
	sell := r.Group("/sell")
	{
		sell.POST("/sell", sellController.Sell)
		sell.GET("/getSellOrder/:id", sellController.GetSellOrder)
		sell.POST("/listSellOrders", sellController.ListSellOrders)
		sell.POST("/updateSellOrder/:id", sellController.UpdateRestockOrder)
		sell.GET("/deleteSellOrder/:id", sellController.DeleteSellOrder)
	}
	stock := r.Group("/stock")
	{
		stock.GET("/get/:id", stockController.Get)
		stock.POST("/list", stockController.List)
		stock.POST("/update/:id", stockController.Update)
	}
	customer := r.Group("/customer")
	{
		customer.POST("/create", customerController.Create)
		customer.GET("/get/:id", customerController.Get)
		customer.POST("/list", customerController.List)
		customer.POST("/update/:id", customerController.Update)
		customer.GET("/delete/:id", customerController.Delete)
	}
	supplier := r.Group("/supplier")
	{
		supplier.POST("/create", supplierController.Create)
		supplier.GET("/get/:id", supplierController.Get)
		supplier.POST("/list", supplierController.List)
		supplier.POST("/update/:id", supplierController.Update)
		supplier.GET("/delete/:id", supplierController.Delete)
	}

	r.Run(global.Conf.Host + ":" + global.Conf.Port)
}
