package service

import (
	"errors"
	"golang_web/dao"
	"golang_web/model"
	"golang_web/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error

	iUser := m.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)

	if iUser.ID == 0 {
		errResult = errors.New("invalid UserName Or Password")
	}
	return iUser, errResult
}
