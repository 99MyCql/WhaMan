package main

import (
	"WhaMan/app/customer"
	"WhaMan/app/restock"
	sellController "WhaMan/app/sell"
	"WhaMan/app/supplier"
	"WhaMan/app/user"
	_ "WhaMan/docs"
	"WhaMan/middleware"
	"WhaMan/pkg/config"
	"WhaMan/pkg/database"
	"WhaMan/pkg/datetime"
	"WhaMan/pkg/log"
	"WhaMan/pkg/validate"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	config.Init("conf.yml")     // 初始化配置
	log.Init(logrus.DebugLevel) // 初始化日志
	database.Init()             // 初始化数据库

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("datetime", validate.DatetimeFormat); err != nil {
			log.Logger.Fatal(err)
		}
		v.RegisterCustomTypeFunc(validate.MyDatetimeValidate, datetime.MyDatetime{})
	}

	r := gin.Default()

	/*** 设置中间件 ***/
	// 设置HTTPS
	r.Use(middleware.TlsHandler())
	// 日志中间件
	r.Use(middleware.Log())
	// 设置基于cookie的session中间件
	store := cookie.NewStore([]byte(config.Conf.SessionSecret))
	r.Use(sessions.Sessions("WhaManSession", store))

	/*** 配置路由 ***/
	// debug模式下注册 swagger 路由
	if config.Conf.Env == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 业务路由
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", user.Login)
	}
	restockRouter := r.Group("/restock")
	restockRouter.Use(middleware.AuthSession())
	{
		restockRouter.POST("/create", restock.Create)
		restockRouter.GET("/get/:id", restock.Get)
		restockRouter.POST("/list", restock.List)
		restockRouter.POST("/listGroupByModelNum", restock.ListGroupByModelNum)
		restockRouter.POST("/update/:id", restock.Update)
		restockRouter.GET("/delete/:id", restock.Delete)
	}
	sellRouter := r.Group("/sell")
	sellRouter.Use(middleware.AuthSession())
	{
		sellRouter.POST("/create", sellController.Create)
		sellRouter.GET("/get/:id", sellController.Get)
		sellRouter.POST("/list", sellController.List)
		sellRouter.POST("/update/:id", sellController.Update)
		sellRouter.GET("/delete/:id", sellController.Delete)
	}
	customerRouter := r.Group("/customer")
	customerRouter.Use(middleware.AuthSession())
	{
		customerRouter.POST("/create", customer.Create)
		customerRouter.GET("/get/:id", customer.Get)
		customerRouter.POST("/list", customer.List)
		customerRouter.POST("/update/:id", customer.Update)
		customerRouter.GET("/delete/:id", customer.Delete)
	}
	supplierRouter := r.Group("/supplier")
	supplierRouter.Use(middleware.AuthSession())
	{
		supplierRouter.POST("/create", supplier.Create)
		supplierRouter.GET("/get/:id", supplier.Get)
		supplierRouter.POST("/list", supplier.List)
		supplierRouter.POST("/update/:id", supplier.Update)
		supplierRouter.GET("/delete/:id", supplier.Delete)
	}

	r.RunTLS(config.Conf.Host+":"+config.Conf.Port, config.Conf.SslCert, config.Conf.SslKey)
}
