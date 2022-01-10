package config

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/uptrace/bun"
)

const CLIENT_URL = "http://localhost:8000"

type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	DB           *bun.DB
	InProduction bool
	Cache        *redis.Client
	TasksClient  *asynq.Client
}
