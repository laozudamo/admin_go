package main

import (
	"admin_go/initialize"

	"go.uber.org/zap"
)

// test git
func main() {
	// 初始化配置文件
	initialize.InitConfig()

	// 初始化数据库
	initialize.InitMysqlDB()

	// 初始化Logger
	initialize.InitLogger()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	Router := initialize.InitRouter()
	err := Router.Run(":8000")
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
	}
}
