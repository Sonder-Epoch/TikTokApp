package main

import (
	"TikTokApp/config"
	"TikTokApp/dao"
	"TikTokApp/logger"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()
	server.Use(logger.GinLogger(), logger.GinRecovery(true))
	setUpRouter(server)
	err := server.Run(config.Conf.Server.Port)
	if err != nil {
		panic(err)
	}
}
func init() {
	dao.InitDB()
	dao.InitRedis()
	logger.InitLogger()
}
