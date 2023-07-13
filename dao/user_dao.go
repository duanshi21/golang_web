package dao

import (
	"golang_web/model"
	"golang_web/service/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBaseDao()}
	}

	return userDao
}

// GetUserByNameAndPassword 查询
func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}

// CheckUserNameExist 查询用户是否存在
func (m *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("name = ?", stUserName).
		Count(&nTotal)
	return nTotal > 0
}

// AddUser 添加用户
func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)

	err := m.Orm.Save(&iUser).Error
	if err != nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}
	return err
}

// GetUserById 根据id查询
func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := m.Orm.First(&iUser, id).Error
	return iUser, err
}
