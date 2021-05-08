package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/yyyThree/gin/config"
	"time"
)

var client *redis.Client

func GetConn() *redis.Client {
	if client != nil && client.Ping(GetCtx()).Err() == nil {
		return client
	}
	client = redis.NewClient(&redis.Options{
		Addr:         config.Config.Redis.Address,
		Password:     config.Config.Redis.Password,
		DB:           config.Config.Redis.DB,
		DialTimeout:  time.Duration(config.Config.Redis.ConnectTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Config.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Config.Redis.WriteTimeout) * time.Second,
		PoolSize:     config.Config.Redis.PoolSize,
	})
	return client
}

func GetCtx() context.Context {
	return context.Background()
}
