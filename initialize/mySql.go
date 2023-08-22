package initialize

import (
	"admin_go/global"
	"admin_go/models"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() {
	mysqlInfo := global.Settings.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host,
		mysqlInfo.Port, mysqlInfo.DBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接池链接设置最大可复用时间
	db.AutoMigrate(models.SysUser{}, models.SysRole{})

	global.DB = db
}
