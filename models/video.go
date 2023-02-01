package models

import (
	"TikTokApp/dao"
	"fmt"
	"sync"
)

type Video struct {
	Id             int64
	UId            int64
	PlayUrl        string
	CoverUrl       string
	CommentCount   int64
	FavouriteCount int64
	Title          string
	CreateTime     int64 `gorm:"autoCreateTime:milli" `
	UpdateTime     int64 `gorm:"autoUpdateTime:milli" `
	IsDeleted      bool
}
type VideoDTO struct {
	Id            int64   `json:"id"`
	Author        UserDTO `json:"author"`
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount int64   `json:"favorite_count"`
	CommentCount  int64   `json:"comment_count"`
	IsFavorite    bool    `json:"is_favorite"`
	Title         string  `json:"title"`
}

type VideoDao struct {
}

var videoDao *VideoDao //DAO(DataAccessObject)模式
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// IsFavourite 查询视频是否已经点赞
func (d VideoDao) IsFavourite(uid, vid int64) bool {
	redisKey := fmt.Sprintf("vifeo:like:%d", vid)
	result, err := dao.REDIS.GetBit(dao.CTX, redisKey, uid).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// FindVideoFeed 查询视频接口流
func (d VideoDao) FindVideoFeed(latest int64, limit int) (videoDTOList []VideoDTO, nextTime int64) {
	var videoList []Video
	dao.DB.Model(&Video{}).
		Order("create_time").
		//Where("create_time>=", latest).
		Limit(limit).Scan(&videoList)
	return createVideoDTOList(&videoList)
}

// 构建视频dto切片
func createVideoDTOList(videoList *[]Video) (videoDTOList []VideoDTO, nextTime int64) {
	nextTime = (*videoList)[0].CreateTime
	for _, video := range *videoList {
		if video.CreateTime < nextTime {
			nextTime = video.CreateTime
		}
		videoDTOList = append(videoDTOList, createVideoDTO(video))
	}
	return
}

// 构建dto
func createVideoDTO(video Video) VideoDTO {
	return VideoDTO{
		Id:            video.Id,
		Author:        userDao.GetUserDTO(video.UId),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavouriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    videoDao.IsFavourite(video.Id, video.UId),
		Title:         video.Title,
	}
}
