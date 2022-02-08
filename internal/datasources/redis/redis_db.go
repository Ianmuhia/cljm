package redisDb

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8" //nolint:goimports
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func GetRedisClient() *redis.Client {
	//TODO:use address from env file
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		//logger.Error("error connecting to redis", err)
		panic(err)

	}
	log.Println("Redis database  connected successfully...")
	return RedisClient
}
