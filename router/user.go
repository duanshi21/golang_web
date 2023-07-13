package router

import (
	"github.com/gin-gonic/gin"
	"golang_web/api"
)

func InitUserRoutes() {
	ResistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {

			}
		}())
		{
			rgPublic.POST("/login", userApi.Login)
		}
		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserById)
		}
	})
}
