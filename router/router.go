package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "golang_web/docs"
)

type IFnRegisterRoute = func(rePublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRegisterRoute
)

// ResistRoute 注册路由
func ResistRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

// InitRouter 初始化系统路由
func InitRouter() {
	// 初始化gin框架，并注册相关路由
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public") // 公开
	rgAuth := r.Group("/api/v1")          // 鉴权

	InitBasePlatformRoutes()

	for _, fnRegisRoute := range gfnRouters {
		fnRegisRoute(rgPublic, rgAuth)
	}

	// 集成Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 从配置文件中读取并配置web服务配置
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8090"
	}

	err := r.Run(":" + stPort)
	if err != nil {
		panic(fmt.Sprintf("Start Server Error：%s", err.Error()))
	}
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
