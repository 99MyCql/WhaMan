package database

import (
	"log"
	"os"
	"time"

	customerDO "WhaMan/app/customer/do"
	restockDO "WhaMan/app/restock/do"
	sellDO "WhaMan/app/sell/do"
	stockDO "WhaMan/app/stock/do"
	supplierDO "WhaMan/app/supplier/do"
	"WhaMan/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库操作对象
var DB *gorm.DB

// Init 初始化数据库连接
func Init() {
	// 创建数据库连接池
	var err error
	DB, err = gorm.Open(mysql.Open(config.Conf.MysqlUrl), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "[WhaMan-DB] ", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // 日志级别
				Colorful:      true,        // 彩色打印
				// IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			},
		),
	})
	if err != nil {
		panic(err)
	}

	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。
	// 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。
	// 出于保护您数据的目的，它不会删除未使用的列
	err = DB.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8").
		AutoMigrate(&customerDO.Customer{}, &stockDO.Stock{}, &supplierDO.Supplier{},
			&restockDO.RestockOrder{}, &sellDO.SellOrder{})
	if err != nil {
		panic(err)
	}
}
