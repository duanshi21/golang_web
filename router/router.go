package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRegisterRoute = func(rePublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRegisterRoute
)

func ResistRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}
func InitRouter() {
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public") // 公开
	rgAuth := r.Group("/api/v1")          // 鉴权

	InitBasePlatformRoutes()

	for _, fnRegisRoute := range gfnRouters {
		fnRegisRoute(rgPublic, rgAuth)
	}

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	err := r.Run(":" + stPort)
	if err != nil {
		panic(fmt.Sprintf("Start Server Error：%s", err.Error()))
	}
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
