package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	DB = Connect()
}

func Connect() *gorm.DB {
	root := "root"
	password := "123456"
	host := "localhost"
	port := 3306
	dbname := "itcast"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", root, password, host, port, dbname)
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return nil
	} else {
		fmt.Println("连接数据库成功")
	}
	return db
}
