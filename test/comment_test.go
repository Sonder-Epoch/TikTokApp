package test

import (
	"TikTokApp/dao"
	"TikTokApp/models"
	"fmt"
	"testing"
)

func TestCreateComment(t *testing.T) {
	dao.InitDB()
	dao.DB.AutoMigrate(&models.Comment{})
	commentDao := models.NewCommentDaoInstance()
	err := commentDao.CreateComment(&models.Comment{
		VId:     1,
		UId:     1,
		Content: "差不多得了",
	})
	if err != nil {
		fmt.Println("Create comment failed:", err)
	}
}
func TestCommentList(t *testing.T) {
	dao.InitDB()
	commentDao := models.NewCommentDaoInstance()
	list := commentDao.GetCommentsByVideoId(1)
	for _, comment := range list {
		fmt.Printf("id:%v\n", comment.Id)
		fmt.Printf("vid:%v\n", comment.VId)
		fmt.Printf("uid:%v\n", comment.UId)
		fmt.Printf("content:%v\n", comment.Content)
		fmt.Printf("update:%v\n", comment.UpdateTime)
		fmt.Printf("create:%v\n", comment.CreateTime)
	}
}
