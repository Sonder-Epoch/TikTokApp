package dao

import (
	"TikTokApp/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct{}

var CTX = context.Background()
var REDIS *redis.Client

func InitRedis() {
	redisConf := config.Conf.Redis
	REDIS = redis.NewClient(&redis.Options{
		Addr:     redisConf.Address,
		DB:       redisConf.DB,
		PoolSize: redisConf.PoolSize,
	})
	_, err := REDIS.Ping(CTX).Result()
	if err != nil {
		panic(fmt.Errorf("连接redis出错，错误信息：%v", err))
	}
}
