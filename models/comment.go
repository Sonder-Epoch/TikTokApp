package models

import (
	"TikTokApp/dao"
	"gorm.io/plugin/soft_delete"
	"sync"
)

type Comment struct {
	Id         int64                 `gorm:"column:id"`
	VId        int64                 `gorm:"column:vid"`
	UId        int64                 `gorm:"column:uid"`
	Content    string                `gorm:"column:content"`
	CreateTime int64                 `gorm:"autoCreateTime:milli" `
	UpdateTime int64                 `gorm:"autoUpdateTime:milli" `
	IsDeleted  soft_delete.DeletedAt `gorm:"column:is_deleted;softDelete:flag"`
}

type CommentDao struct {
}

var commentDao *CommentDao //DAO(DataAccessObject)模式
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

// CreateComment 创建评论
func (*CommentDao) CreateComment(comment *Comment) error {
	return dao.DB.Create(comment).Error
}

// GetCommentsByVideoId 根据视频id查询评论
func (*CommentDao) GetCommentsByVideoId(vid int64) (commentList []Comment) {
	dao.DB.Model(&Comment{}).Where("vid=?", vid).Order("create_time").Find(&commentList)
	return
}

// RemoveComment 删除评论
func (*CommentDao) RemoveComment(cid string) error {
	return dao.DB.Delete(&Comment{}, cid).Error
}

func (*CommentDao) GetCommentById(cid int64) (comment Comment) {
	dao.DB.Model(&Comment{}).Where("id = ?", cid).Scan(&comment)
	return
}
