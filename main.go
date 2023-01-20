package main

import (
	"TikTokApp/config"
	"TikTokApp/dao"
	"TikTokApp/logger"
	"TikTokApp/utils"
)

func main() {
	server := setUpRouter()
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
