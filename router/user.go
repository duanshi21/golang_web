package router

import (
	"github.com/gin-gonic/gin"
	"golang_web/api"
	"net/http"
)

func InitUserRoutes() {
	ResistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublic = rgPublic.Group("user")
		{
			rgPublic.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "Jack"},
					{"id": 2, "name": "Tom"},
				},
			})
		})
		rgAuthUser.GET("/:id", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "Tom",
			})
		})
	})
}
