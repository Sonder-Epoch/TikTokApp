package models

import (
	"TikTokApp/dao"
	"fmt"
	"sync"
	"time"
)

type Video struct {
	Id             int64     `gorm:"column:id" redis:"id"`
	UId            int64     `gorm:"column:uid" redis:"uid"`
	PlayUrl        string    `gorm:"column:play_url" redis:"play_url"`
	CoverUrl       string    `gorm:"column:cover_url" redis:"cover_url"`
	CommentCount   int64     `gorm:"column:comment_count" redis:"comment_count"`
	FavouriteCount int64     `gorm:"column:favourite_count" redis:"favorite_count"`
	Title          string    `gorm:"column:title" redis:"title"`
	CreateTime     time.Time `gorm:"column:create_time" redis:"-"`
	UpdateTime     time.Time `gorm:"column:update_time" redis:"-"`
	IsDeleted      bool      `gorm:"column:is_deleted" redis:"-"`
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
func (d *VideoDao) IsFavourite(uid, vid int64) bool {
	redisKey := fmt.Sprintf("vifeo:like:%d", vid)
	result, err := dao.REDIS.GetBit(dao.CTX, redisKey, uid).Result()
	if err != nil {
		return false
	}
	return result == 1
}
