package test

import (
	"TikTokApp/dao"
	"TikTokApp/models"
	"fmt"
	"testing"
)

func TestCreateVideo(t *testing.T) {
	dao.InitDB()
	dao.DB.AutoMigrate(&models.Video{})
	dao.DB.AutoMigrate(&models.Favorite{})
	videoDao := models.NewVideoDaoInstance()
	err := videoDao.CreateVideo(&models.Video{
		Id:             1,
		AuthorId:       1,
		PlayUrl:        "test",
		CoverUrl:       "Test",
		CommentCount:   1,
		FavouriteCount: 1,
		Title:          "测试",
	})
	if err != nil {
		fmt.Println("Create video failed:", err)
	}
}
func TestFavoriteList(t *testing.T) {
	dao.InitDB()
	videoDao := models.NewVideoDaoInstance()
	list := videoDao.FavoriteList(1)
	for _, video := range list {
		fmt.Println(video)
	}
}

func TestFavoriteVideo(t *testing.T) {
	dao.InitDB()
	dao.InitRedis()
	videoDao := models.NewVideoDaoInstance()
	err := videoDao.FavoriteVideo(1, 1, 1)
	if err != nil {
		fmt.Println("Favorite video failed:", err)
	}
}
func TestIsFavorite(t *testing.T) {
	dao.InitDB()
	dao.InitRedis()
	videoDao := models.NewVideoDaoInstance()
	fmt.Println(videoDao.IsFavourite(1, 1))
}
func TestGetVideoList(t *testing.T) {
	dao.InitDB()
	videoDao := models.NewVideoDaoInstance()
	for i, video := range videoDao.GetVideoList(0, 30) {
		fmt.Printf("video_index:   %d\n", i)
		fmt.Printf("Id            :%v\n", video.Id)
		fmt.Printf("AuthorId      :%v\n", video.AuthorId)
		fmt.Printf("PlayUrl       :%v\n", video.PlayUrl)
		fmt.Printf("CoverUrl      :%v\n", video.CoverUrl)
		fmt.Printf("CommentCount  :%v\n", video.CommentCount)
		fmt.Printf("FavouriteCount:%v\n", video.FavouriteCount)
		fmt.Printf("Title         :%v\n", video.Title)
		fmt.Printf("CreateTime    :%v\n", video.CreateTime)
		fmt.Printf("UpdateTime    :%v\n", video.UpdateTime)
		fmt.Printf("IsDeleted     :%v\n", video.IsDeleted)
	}
}
