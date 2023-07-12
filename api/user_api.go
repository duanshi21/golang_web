package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang_web/service/dto"
	"golang_web/utils"
	"reflect"
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
	var iUserLoginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&iUserLoginDTO)
	if errs != nil {
		Fail(ctx, ResponseJson{
			Msg: parseValidateErrors(errs.(validator.ValidationErrors), &iUserLoginDTO).Error(),
		})
		return
	}

	OK(ctx, ResponseJson{
		Data: iUserLoginDTO,
	})

	//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})
	//ok(ctx, ResponseJson{})
	//ResponseOK(ctx, Response{
	//	Msg: "Login Success",
	//})
	//
	//ResponseError(ctx, Response{
	//	Msg: "Login Failure",
	//})
	//OK(ctx, ResponseJson{
	//	Msg: "Login Failure",
	//})
}

func parseValidateErrors(errs validator.ValidationErrors, target any) error {
	var errResult error

	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errs {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s:%s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))

	}
	return errResult
}
