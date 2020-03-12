package main

import (
	"Service/logger"
	"Service/init"
	"Service/web"
	"github.com/spf13/viper"
)

func main() {
	logger.LOG_INFO("Service Starting!")
	//加载配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.ReadInConfig()

	//初始化模块
	initdata.Init()

	//开启http监听
	http_port := viper.GetString("params.port")
	web.Start(http_port)
}
