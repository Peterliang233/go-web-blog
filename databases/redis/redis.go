package redis

import (
	"context"
	"fmt"
	"github.com/Peterliang233/go-blog/configs"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisClient *redis.Client

// InitRedis 初始化redis的操作
func InitRedis() {
	redisAddr := fmt.Sprintf("%s:%s", configs.RdHost, configs.RdPort)

	RedisClient = redis.NewClient(
		&redis.Options{
			Addr:     redisAddr,
			Password: configs.RdPassword,
			DB:       0,
		},
	)

	pong, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("redis 启动错误 %v %v\n", pong, err)
	}

	log.Println("redis start success")
}
