package redis_db

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)
var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func init() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "172.17.0.3:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping(Ctx).Result()

	fmt.Println(pong, err)
}

func GetRedisClient() *redis.Client {
	return RedisClient
}
