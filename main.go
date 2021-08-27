package main

import (
	customerController "WhaMan/app/customer/controller"
	restockController "WhaMan/app/restock/controller"
	sellController "WhaMan/app/sell/controller"
	stockController "WhaMan/app/stock/controller"
	supplierController "WhaMan/app/supplier/controller"
	userController "WhaMan/app/user/controller"
	_ "WhaMan/docs"
	"WhaMan/middleware"
	"WhaMan/pkg/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/global"
	"WhaMan/pkg/log"
	"WhaMan/pkg/validators"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("datetime", validators.DatetimeFormat); err != nil {
			global.Log.Fatal(err)
		}
	}

	r := gin.Default()

	// 设置中间件
	// 设置HTTPS
	r.Use(middleware.TlsHandler())
	// 设置基于cookie的session中间件
	store := cookie.NewStore([]byte(global.Conf.SessionSecret))
	r.Use(sessions.Sessions("WhaManSession", store))

	// 配置路由
	// debug模式下注册 swagger 路由
	if global.Conf.Env == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 业务路由
	user := r.Group("/user")
	{
		user.POST("/login", userController.Login)
	}
	restock := r.Group("/restock")
	restock.Use(middleware.AuthSession())
	{
		restock.POST("/restock", restockController.Restock)
		restock.GET("/getRestockOrder/:id", restockController.GetRestockOrder)
		restock.POST("/listRestockOrders", restockController.ListRestockOrders)
		restock.POST("/updateRestockOrder/:id", restockController.UpdateRestockOrder)
		restock.GET("/deleteRestockOrder/:id", restockController.DeleteRestockOrder)
	}
	sell := r.Group("/sell")
	sell.Use(middleware.AuthSession())
	{
		sell.POST("/sell", sellController.Sell)
		sell.GET("/getSellOrder/:id", sellController.GetSellOrder)
		sell.POST("/listSellOrders", sellController.ListSellOrders)
		sell.POST("/updateSellOrder/:id", sellController.UpdateRestockOrder)
		sell.GET("/deleteSellOrder/:id", sellController.DeleteSellOrder)
	}
	stock := r.Group("/stock")
	stock.Use(middleware.AuthSession())
	{
		stock.GET("/get/:id", stockController.Get)
		stock.POST("/list", stockController.List)
		stock.POST("/update/:id", stockController.Update)
	}
	customer := r.Group("/customer")
	customer.Use(middleware.AuthSession())
	{
		customer.POST("/create", customerController.Create)
		customer.GET("/get/:id", customerController.Get)
		customer.POST("/list", customerController.List)
		customer.POST("/update/:id", customerController.Update)
		customer.GET("/delete/:id", customerController.Delete)
	}
	supplier := r.Group("/supplier")
	supplier.Use(middleware.AuthSession())
	{
		supplier.POST("/create", supplierController.Create)
		supplier.GET("/get/:id", supplierController.Get)
		supplier.POST("/list", supplierController.List)
		supplier.POST("/update/:id", supplierController.Update)
		supplier.GET("/delete/:id", supplierController.Delete)
	}

	// r.Run(global.Conf.Host + ":" + global.Conf.Port)
	r.RunTLS(global.Conf.Host+":"+global.Conf.Port, global.Conf.SslCert, global.Conf.SslKey)
}
