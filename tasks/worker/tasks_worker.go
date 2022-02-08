package main

import (
	"log"
	"maranatha_web/internal/services"

	"github.com/hibiken/asynq" //nolint:goimports
)

const (
	TypeWelcomeEmail = "email:welcome"
	TypeDailyVerse   = "dailyVerse"
)

// workers.go
func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency:    10,
			RetryDelayFunc: nil,
			IsFailure:      nil,
			Queues:         nil,
			StrictPriority: false,
			ErrorHandler:   nil,
			//Logger:              zap.,
			LogLevel:            0,
			ShutdownTimeout:     0,
			HealthCheckFunc:     nil,
			HealthCheckInterval: 0,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeWelcomeEmail, services.HandleVerifyEmailTask)
	mux.HandleFunc(TypeDailyVerse, services.HandleDailyVerseTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
