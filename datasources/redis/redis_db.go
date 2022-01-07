package redis_db

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"maranatha_web/logger"
)

var (
	RedisClient *redis.Client
)
var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.Background()
)

func init() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		logger.Error("error connecting to redis", err)
		panic(err)

	}
	fmt.Println(pong, err)
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
