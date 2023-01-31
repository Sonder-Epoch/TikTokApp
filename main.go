package main

import (
	"TikTokApp/config"
	"TikTokApp/dao"
	"TikTokApp/logger"
	"TikTokApp/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	setUpRouter(server)
	if err := server.Run(config.Conf.Server.Port); err != nil {
		panic(err)
	}
}
func init() {
	dao.InitDB()
	dao.InitRedis()
	logger.InitLogger()
	utils.InitOSS()
}
