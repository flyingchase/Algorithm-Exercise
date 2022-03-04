package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb   *redis.Client
	ctx   = context.Background()
	mutex sync.Mutex
)

func NewRedis() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:7890",
			Password: "",
			DB:       0,
		})
}
func Lock(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	bool, err := rdb.SetNX(ctx, key, 1, 10*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return bool
}
func UnLock(key string) int64 {
	nums, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
