package db

import (
	"LearningGo/configs"
	"LearningGo/log"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		configs.DbSettings.Root,
		configs.DbSettings.Password,
		configs.DbSettings.Host,
		configs.DbSettings.Port,
		configs.DbSettings.Dbname,
		configs.DbSettings.Charset,
		configs.DbSettings.ParseTime,
		configs.DbSettings.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.SugarLogger.Error(err)
		return nil
	}
	fmt.Println("连接数据库成功")
	return db
}
