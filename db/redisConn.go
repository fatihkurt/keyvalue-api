package db

import (
	"context"
	"deliveryhero/helper"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx    context.Context
	client *redis.Client
	once   sync.Once
)

func RedisClient() *redis.Client {
	var err error

	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     helper.GetEnv("REDIS_URL", "localhost:6379"),
			Password: helper.GetEnv("REDIS_PASSWORD", ""),
			DB:       0,
		})
	})

	if err != nil {
		panic(err)
	}

	return client
}
