package config

import (
	"go.uber.org/zap"
)

const CLIENT_URL = "http://localhost:8000"

type AppConfig struct {
	InfoLog  *zap.Logger
	ErrorLog *zap.Logger
	// MailServer *mail.SMTPServer
	// MailClient *mail.SMTPClient
	// DB           *gorm.DB
	// RedisClient  *redis.Client
	InProduction bool
	// Cache        *redis.Client
	// TasksClient  *asynq.Client
}

func NewAppConfig(infoLog *zap.Logger, errorLog *zap.Logger) *AppConfig {
	return &AppConfig{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

}
