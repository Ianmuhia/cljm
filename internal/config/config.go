package config

import (
	"net/url"

	"go.uber.org/zap"
)

type AppConfig struct {
	InfoLog                 *zap.Logger
	ErrorLog                *zap.Logger
	StorageURL              *url.URL
	StorageBucket           string
	PasswordResetCodeExpiry int
	InProduction            bool
	TokenLifeTime           int
}

func NewAppConfig(infoLog *zap.Logger, errorLog *zap.Logger) *AppConfig {
	return &AppConfig{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

}
