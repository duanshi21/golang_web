package cmd

import (
	"fmt"
	"golang_web/conf"
	"golang_web/global"
	"golang_web/router"
	"golang_web/utils"
)

func Start() {
	var initErr error

	fmt.Println("========Start========")
	// 初始化系统配置文件
	conf.InitConfig()
	// 初始化日志组件
	global.Logger = conf.InitLogger()
	// 初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// 初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean========")
}
