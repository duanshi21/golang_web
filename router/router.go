package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "golang_web/docs"
	"strings"
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

	// 初始化平台的路由
	InitBasePlatformRoutes()

	// 注册自定义验证器
	registryCostValidator()

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

func registryCostValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && 0 == strings.Index(value, "a") {
					return true
				}
			}
			return false
		})
	}
}
