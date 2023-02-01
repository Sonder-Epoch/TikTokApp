package models

import (
	"TikTokApp/dao"
	"go.uber.org/zap"
	"sync"
	"time"
)

type User struct {
	Id            int64     `gorm:"column:id" redis:"id"`
	Name          string    `gorm:"column:name" redis:"name"`
	Password      string    `gorm:"column:password" redis:"password"`
	FollowCount   int64     `gorm:"column:follow_count" redis:"follow_count"`
	FollowerCount int64     `gorm:"column:follower_count" redis:"follower_count"`
	CreateTime    time.Time `gorm:"column:create_time" redis:"-"`
	UpdateTime    time.Time `gorm:"column:update_time" redis:"-"`
	IsDeleted     bool      `gorm:"column:is_deleted" redis:"-"`
}
type UserDTO struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}
func (*UserDao) CreateUser(user *User) {
	if err := dao.DB.Create(user).Error; err != nil {
		zap.L().Sugar().Errorf("创建user失败user")
	}
}

func (d UserDao) GetUserDTO(id int64) UserDTO {

}
