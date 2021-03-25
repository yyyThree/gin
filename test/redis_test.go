package test

import (
	"fmt"
	"gin/config"
	"gin/library/redis"
	"testing"
)

var key = "redis"

func init() {
	config.Load()
}

func TestSet(t *testing.T) {
	res, err := redis.GetConn().Set(redis.GetCtx(), key, 1, 0).Result()
	if err != nil {
		t.Fatal("redis设置失败", err)
	}
	fmt.Println("redis设置成功\n", res)
}

func TestGet(t *testing.T) {
	res, err := redis.GetConn().Get(redis.GetCtx(), key).Result()
	if err != nil {
		t.Fatal("redis获取失败", err)
	}
	fmt.Println("redis获取成功\n", res)
}