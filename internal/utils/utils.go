package utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"time"

	redis_db "maranatha_web/internal/datasources/redis"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateRandomExpiryCode(key string) string {

	dura := 10 * time.Minute

	var d = dura
	fmt.Printf("Redis TTL %v", d.Minutes())
	max := 6
	value := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, value, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(value); i++ {
		value[i] = table[int(value[i])%len(table)]
	}

	redis_db.RedisClient.Set(context.TODO(), key, value, d)

	log.Println(string(value))
	return string(value)
}
