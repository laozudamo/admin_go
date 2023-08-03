package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 设置配置文件路径
	config := "config.yaml"

	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}
	// 初始化 viper 配置
	v := viper.New()
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	// 监听配置文件 并更新

}
