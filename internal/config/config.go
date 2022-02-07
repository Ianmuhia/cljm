package config

import (
	"net/url"

	"go.uber.org/zap"
)

const CLIENT_URL = "http://localhost:8000"

type AppConfig struct {
	InfoLog                 *zap.Logger
	ErrorLog                *zap.Logger
	StorageURL              *url.URL
	StorageBucket           string
	PasswordResetCodeExpiry int
	InProduction            bool
}

func NewAppConfig(infoLog *zap.Logger, errorLog *zap.Logger) *AppConfig {
	return &AppConfig{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

}
