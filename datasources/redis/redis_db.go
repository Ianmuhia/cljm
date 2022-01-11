package redis_db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8" //nolint:goimports
	"maranatha_web/logger"         //nolint:goimports
)

var (
	RedisClient *redis.Client
)
var (
	Ctx = context.Background()
)

func GetRedisClient() *redis.Client {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		logger.Error("error connecting to redis", err)
		panic(err)

	}
	log.Println("Redis database  connected successfully...")
	return RedisClient
}
