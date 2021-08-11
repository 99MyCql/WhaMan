package database

import (
	customerModel "WhaMan/app/customer/model"
	restockModel "WhaMan/app/restock/model"
	sellModel "WhaMan/app/sell/model"
	stockModel "WhaMan/app/stock/model"
	supplierModel "WhaMan/app/supplier/model"
	"WhaMan/pkg/global"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New 初始化数据库连接
func New() (*gorm.DB, error) {
	// 创建数据库连接池
	db, err := gorm.Open(mysql.Open(global.Conf.MysqlUrl), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "连接数据库失败")
	}

	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。
	// 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。
	// 出于保护您数据的目的，它不会删除未使用的列
	err = db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8").
		AutoMigrate(&customerModel.Customer{}, &stockModel.Stock{}, &supplierModel.Supplier{},
			&restockModel.RestockOrder{}, &sellModel.SellOrder{})
	if err != nil {
		return nil, errors.Wrap(err, "数据库自动迁移失败")
	}

	return db, nil
}
