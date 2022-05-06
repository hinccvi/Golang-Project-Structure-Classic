package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hinccvi/go-template/log"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
		PoolSize: 0,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Panic("Fail to connect to Redis")
	} else {
		log.Info("Redis successfully init")
	}
}
