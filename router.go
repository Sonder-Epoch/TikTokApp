package main

import (
	"TikTokApp/controller"
	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	//server := gin.New()
	////日志自定义
	//server.Use(logger.GinLogger(), logger.GinRecovery(true))
	server := gin.Default()
	//路由配置
	server.GET("/test", controller.Test)
	server.POST("publish", controller.Publish)
	return server
}
