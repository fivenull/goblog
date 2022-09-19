package model

import (
	"fmt"
	"goblog/pkg/logger"

	"gorm.io/gorm"

	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:secret@tcp(10.13.0.252:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	fmt.Println(*config)
	DB, err = gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
}
