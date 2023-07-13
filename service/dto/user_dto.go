package dto

import "golang_web/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,email,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

// UserAddDTO 添加用户相关dto
type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" from:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" from:"real_name"`
	Avatar   string
	Mobile   string `json:"mobile" from:"mobile"`
	Email    string `json:"email" from:"email"`
	Password string `json:"password,omitempty" from:"password" binding:"required" message:"密码不能为空"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
}
