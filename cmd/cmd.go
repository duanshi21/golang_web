package cmd

import (
	"fmt"
	"golang_web/conf"
	"golang_web/global"
	"golang_web/router"
)

func Start() {
	fmt.Println("========Start========")
	// 初始化系统配置文件
	conf.InitConfig()
	// 初始化日志组件
	global.Logger = conf.InitLogger()
	// 初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean========")
}
