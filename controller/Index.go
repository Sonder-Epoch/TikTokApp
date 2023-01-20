package controller

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "你是懂的",
	})
}
