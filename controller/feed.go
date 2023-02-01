package controller

import (
	"TikTokApp/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	latest_time := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(latest_time, 10, 64)
	if err != nil {
		latestTime = time.Now().UnixMilli()
	}
	videoList, nextTime := models.NewVideoDaoInstance().FindVideoFeed(latestTime, 30)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"next_time":   nextTime,
		"video_list":  videoList,
	})
}
