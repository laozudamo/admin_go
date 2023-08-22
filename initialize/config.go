package initialize

import (
	"admin_go/config"
	"admin_go/global"
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	v.SetConfigFile("./settings-dev.yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Print(serverConfig)

	global.Settings = serverConfig

}
