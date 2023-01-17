package main

import (
	"TikTokApp/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	server := gin.Default()
	setUpRouter(server)
	err := server.Run(config.Conf.Port)
	if err != nil {
		panic(err)
	}
}
