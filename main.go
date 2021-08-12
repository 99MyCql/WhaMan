package main

import (
	customerController "WhaMan/app/customer/controller"
	restockController "WhaMan/app/restock/controller"
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
			"message": "hello world",
		})
	})

	// 注册 swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置路由
	v1 := r.Group("/api/v1")
	{
		customer := v1.Group("/customer")
		{
			customer.POST("/create", customerController.Create)
			customer.GET("/get")
			customer.POST("/list")
			customer.GET("/listSellOrders")
			customer.POST("/update")
			customer.GET("/delete")
		}
		restock := v1.Group("/restock")
		{
			restockOrder := restock.Group("/order")
			{
				restockOrder.POST("/create", restockController.Create)
				restockOrder.GET("/get")
				restockOrder.POST("/list")
				restockOrder.POST("/update")
				restockOrder.GET("/delete")
			}
		}
		sell := v1.Group("/sell")
		{
			sell.POST("/create")
			sell.GET("/get")
			sell.POST("/list")
			sell.POST("/update")
			sell.GET("/delete")
		}
		stock := v1.Group("/stock")
		{
			stock.POST("/create")
			stock.GET("/get")
			stock.POST("/list")
			stock.POST("/update")
			stock.GET("/delete")
		}
		supplier := v1.Group("/supplier")
		{
			supplier.POST("/create")
			supplier.GET("/get")
			supplier.POST("/list")
			supplier.GET("/listRestockOrders")
			supplier.POST("/update")
			supplier.GET("/delete")
		}
	}

	r.Run(global.Conf.Host + ":" + global.Conf.Port)
}
