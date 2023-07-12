package api

import (
	"github.com/gin-gonic/gin"
	"golang_web/service/dto"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
	}
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
func (m UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	m.OK(ResponseJson{
		Data: iUserLoginDTO,
	})

	//c.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})
	//ok(c, ResponseJson{})
	//ResponseOK(c, Response{
	//	Msg: "Login Success",
	//})
	//
	//ResponseError(c, Response{
	//	Msg: "Login Failure",
	//})
	//OK(c, ResponseJson{
	//	Msg: "Login Failure",
	//})
}
