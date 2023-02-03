package models

import (
	"TikTokApp/dao"
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Id            int64  `gorm:"column:id;primaryKey" redis:"id"`
	Username      string `gorm:"column:username" redis:"username"`
	Password      string `gorm:"column:password" redis:"password"`
	FollowCount   int64  `gorm:"column:follow_count" redis:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" redis:"follower_count"`
	// CreateTime    time.Time `gorm:"column:create_time" redis:"-"`
	// UpdateTime    time.Time `gorm:"column:update_time" redis:"-"`
	IsDeleted soft_delete.DeletedAt `gorm:"column:is_deleted;softDelete:flag" redis:"-"`
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

// 添加用户
func (*UserDao) CreateUser(user *User) error {
	if err := dao.DB.Create(user).Error; err != nil {
		zap.L().Sugar().Errorf("创建用户失败")
		return err
	}
	return nil
}

// 通过id获取用户信息
func (*UserDao) GetUserById(id int64) (*User, error) {
	user := &User{}
	if err := dao.DB.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Sugar().Errorf("未查询到用户信息")
		} else {
			zap.L().Sugar().Errorf("查询用户失败")
		}
		return user, err
	}
	return user, nil
}

// 通过username和password获取用户信息
func (*UserDao) GetUserByUsernameAndPassword(username string, password string) (*User, error) {
	user := &User{}
	if err := dao.DB.Where("username = ? AND password = ?", username, password).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Sugar().Errorf("未查询到用户信息")
		} else {
			zap.L().Sugar().Errorf("查询用户失败")
		}
		return user, err
	}
	return user, nil
}

// TODO尚未开发
func (*UserDao) IsFollow(uid, authorId int64) bool {
	redisKey := fmt.Sprintf("userfollow:%d", authorId)
	result, err := dao.REDIS.GetBit(dao.CTX, redisKey, uid).Result()
	if err != nil {
		return false
	}
	return result == 1
}
