package controller

import (
	"TikTokApp/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Publish(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		zap.L().Sugar().Errorf("视频传输失败%v", err)
	}
	title := c.PostForm("title")
	_ = service.Publish(title, file)
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "string",
	})
}

func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		zap.L().Sugar().Errorf("查询失败错误的userid")
	}
	publishList := service.PublishList(uid)
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"video_list":  publishList,
	})
}
