package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// Login 用户登录
// @Tag 用户管理
// @Summary 用户登录
// @Description 用登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} string "登录成功"
// @Failure 401 {object} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "Login Success",
	})
}
