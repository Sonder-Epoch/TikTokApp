package dao

import (
	"TikTokApp/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	mysqlConfig := config.Conf.Mysql
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	username := mysqlConfig.Username
	pwd := mysqlConfig.Pwd
	database := mysqlConfig.Database
	charset := mysqlConfig.Charset
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset%s&parseTime=true",
		username, pwd, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	// 数据库连接池配置
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpen)
	DB = db
}

var (
	DB *gorm.DB
)
