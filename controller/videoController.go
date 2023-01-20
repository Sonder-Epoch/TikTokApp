package controller

import (
	"TikTokApp/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Publish(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		zap.L().Sugar().Errorf("视频传输失败%v", err)
		return
	}
	title := c.PostForm("title")
	service.Publish(title, file)
	c.JSON(200, gin.H{})
}
