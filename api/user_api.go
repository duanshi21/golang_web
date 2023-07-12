package api

import (
	"github.com/gin-gonic/gin"
	"golang_web/service"
	"golang_web/service/dto"
	"golang_web/utils"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
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
	// 收集客户端提交过来的数据，放到controller中
	var iUserLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	// 调用controller中具体的业务service
	iUser, err := m.Service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}
	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)
	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
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
