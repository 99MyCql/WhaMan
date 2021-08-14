package main

import (
	customerController "WhaMan/app/customer/controller"
	restockController "WhaMan/app/restock/controller"
	supplierController "WhaMan/app/supplier/controller"
	_ "WhaMan/docs"
	"WhaMan/pkg/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/global"
	"WhaMan/pkg/log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	// 注册 swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置路由
	restock := r.Group("/restock")
	{
		restock.POST("/restock", restockController.Restock)
		restock.GET("/getRestockOrder")
		restock.POST("/listRestockOrders")
		restock.POST("/updateRestockOrder")
		restock.GET("/deleteRestockOrder")
	}
	sell := r.Group("/sell")
	{
		sell.POST("/sell")
		sell.GET("/getSellOrder")
		sell.POST("/listSellOrders")
		sell.POST("/updateSellOrder")
		sell.GET("/deleteSellOrder")
	}
	stock := r.Group("/stock")
	{
		stock.GET("/get")
		stock.POST("/list")
		stock.POST("/update")
		stock.GET("/delete")
	}
	customer := r.Group("/customer")
	{
		customer.POST("/create", customerController.Create)
		customer.GET("/get")
		customer.POST("/list")
		customer.POST("/update")
		customer.GET("/delete")
	}
	supplier := r.Group("/supplier")
	{
		supplier.POST("/create", supplierController.Create)
		supplier.GET("/get")
		supplier.POST("/list")
		supplier.POST("/update")
		supplier.GET("/delete")
	}

	r.Run(global.Conf.Host + ":" + global.Conf.Port)
}
