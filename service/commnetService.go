package service

import (
	"TikTokApp/common"
	"TikTokApp/models"
	"time"
)

var commentDao = models.NewCommentDaoInstance()

// AddComment 添加评论
func AddComment(comment *models.Comment) {
	commentDao.CreateComment(comment)
}

// CommentList 评论列表
func CommentList(vid int64) []common.CommentDTO {
	commentList := commentDao.GetCommentsByVideoId(vid)
	return createCommentDTOList(&commentList)
}
func DeleteComment(cid string) {
	commentDao.RemoveComment(cid)
}

// GetCommentDTO id 查询评论
func GetCommentDTO(cid int64) common.CommentDTO {
	comment := commentDao.GetCommentById(cid)
	return createCommentDTO(&comment)
}

func createCommentDTOList(commentList *[]models.Comment) (commentDTOList []common.CommentDTO) {
	for _, comment := range *commentList {
		commentDTOList = append(commentDTOList, createCommentDTO(&comment))
	}
	return
}

func createCommentDTO(comment *models.Comment) common.CommentDTO {
	userDTO, _ := FindUserById(comment.UId)
	return common.CommentDTO{
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CreateTime, 0).Format("2006-01-02"),
		User:       *userDTO,
	}
}
