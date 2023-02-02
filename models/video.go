package models

import (
	"TikTokApp/common"
	"TikTokApp/dao"
	"fmt"
	"sync"

	"gorm.io/plugin/soft_delete"
)

type Video struct {
	Id             int64                 `gorm:"column:id"`
	AuthorId       int64                 `gorm:"column:author_id"`
	PlayUrl        string                `gorm:"column:play_url"`
	CoverUrl       string                `gorm:"column:cover_url"`
	CommentCount   int64                 `gorm:"column:comment_count"`
	FavouriteCount int64                 `gorm:"column:favorite_count"`
	Title          string                `gorm:"column:title"`
	CreateTime     int64                 `gorm:"autoCreateTime:milli" `
	UpdateTime     int64                 `gorm:"autoUpdateTime:milli" `
	IsDeleted      soft_delete.DeletedAt `gorm:"column:is_deleted;softDelete:flag"`
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
func (d VideoDao) FindVideoFeed(latest int64, limit int) (videoDTOList []common.VideoDTO, nextTime int64) {
	var videoList []Video
	dao.DB.Model(&Video{}).
		Order("create_time").
		//Where("create_time>=", latest).
		Limit(limit).Scan(&videoList)
	return createVideoDTOList(&videoList)
}

// 构建视频dto切片
func createVideoDTOList(videoList *[]Video) (videoDTOList []common.VideoDTO, nextTime int64) {
	nextTime = (*videoList)[0].CreateTime
	for _, video := range *videoList {
		if video.CreateTime < nextTime {
			nextTime = video.CreateTime
		}
		videoDTOList = append(videoDTOList, GetVideoDTO(video))
	}
	return
}

// 构建dto
func GetVideoDTO(video Video) common.VideoDTO {
	return common.VideoDTO{
		Id:            video.Id,
		Author:        userDao.GetUserDTO(video.AuthorId),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavouriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    videoDao.IsFavourite(video.Id, video.AuthorId),
		Title:         video.Title,
	}
}
