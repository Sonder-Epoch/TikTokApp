package test

import (
	"TikTokApp/dao"
	"TikTokApp/models"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	dao.InitDB()
	dao.DB.AutoMigrate(&models.User{})
	userDao := models.NewUserDaoInstance()
	user := &models.User{
		Id:            1,
		Username:      "张三",
		Password:      "123456",
		FollowCount:   20,
		FollowerCount: 30,
		IsDeleted:     0,
	}
	if err := userDao.CreateUser(user); err != nil {
		fmt.Print("CreateUser Failure!\n")
	}
}

func TestGetUserById(t *testing.T) {
	dao.InitDB()
	userDao := models.NewUserDaoInstance()
	user, err := userDao.GetUserById(1)
	if err != nil {
		fmt.Print("Record Not Found!\n")
	} else {
		fmt.Printf("Id:%v\n", user.Id)
		fmt.Printf("Username:%v\n", user.Username)
		fmt.Printf("Password:%v\n", user.Password)
		fmt.Printf("FollowCount:%v\n", user.FollowCount)
		fmt.Printf("FollowerCount:%v\n", user.FollowerCount)
		// fmt.Printf("IsDeleted:%v\n", user.IsDeleted)
	}
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	dao.InitDB()
	userDao := models.NewUserDaoInstance()
	user, err := userDao.GetUserByUsernameAndPassword("张三", "123456")
	if err != nil {
		fmt.Print("Record Not Found!\n")
	} else {
		fmt.Printf("Id:%v\n", user.Id)
		fmt.Printf("Username:%v\n", user.Username)
		fmt.Printf("Password:%v\n", user.Password)
		fmt.Printf("FollowCount:%v\n", user.FollowCount)
		fmt.Printf("FollowerCount:%v\n", user.FollowerCount)
		// fmt.Printf("IsDeleted:%v\n", user.IsDeleted)
	}
}
