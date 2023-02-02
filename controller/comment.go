package controller

import (
	"TikTokApp/common"
	"TikTokApp/models"
	"TikTokApp/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	if actionType == "1" {
		vid, _ := strconv.ParseInt(videoId, 0, 64)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "string",
			"comment":     deleteComment(c.Query("comment_text"), vid),
		})
	} else if actionType == "2" {
		commentId := c.Query("comment_id")
		service.DeleteComment(commentId)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "string",
		})
	}
}

func deleteComment(commentText string, vid int64) common.CommentDTO {
	comment := models.Comment{
		UId:     common.GetUser(),
		Content: commentText,
		VId:     vid}
	service.AddComment(&comment)
	return service.GetCommentDTO(comment.Id)
}

func CommentList(c *gin.Context) {
	videoId := c.Query("video_id")
	vid, _ := strconv.ParseInt(videoId, 0, 64)
	commentList := service.CommentList(vid)
	c.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   "string",
		"comment_list": commentList,
	})
}
