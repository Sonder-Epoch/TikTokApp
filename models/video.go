package models

import (
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
func (*VideoDao) IsFavourite(uid, vid int64) bool {
	redisKey := fmt.Sprintf("vifeo:like:%d", vid)
	result, err := dao.REDIS.GetBit(dao.CTX, redisKey, uid).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// GetVideoList 获取视频列表
func (*VideoDao) GetVideoList(latest int64, limit int) (videoList []Video) {
	dao.DB.Model(&Video{}).
		Order("create_time").
		//Where("create_time>=", latest).
		Limit(limit).Scan(&videoList)
	return
}

// CreateVideo 创建视频记录
func (*VideoDao) CreateVideo(video *Video) error {
	return dao.DB.Create(video).Error
}

// GetVideoListByUid 根据用户查询视频
func (*VideoDao) GetVideoListByUid(uid int64) (videoList []Video) {
	dao.DB.Model(&Video{}).
		Order("create_time").
		Where("author_id=?", uid).
		Scan(&videoList)
	return
}
