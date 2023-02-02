package service

import (
	"TikTokApp/common"
	"TikTokApp/models"
)

var userDao *models.UserDao = models.NewUserDaoInstance()

func AddUser(user *models.User) error {
	return userDao.CreateUser(user)
}

func FindUserById(id int64) (*common.UserDTO, error) {
	user, err := userDao.GetUserById(id)
	userDto := GetUserDTO(user)
	return userDto, err
}

func FindUserByUsernameAndPassword(username string, password string) (*common.UserDTO, error) {
	user, err := userDao.GetUserByUsernameAndPassword(username, password)
	userDto := GetUserDTO(user)
	return userDto, err
}

func GetUserDTO(user *models.User) *common.UserDTO {
	return &common.UserDTO{
		Id:            user.Id,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      userDao.IsFollow(1), //该接口尚未开发
	}
}
