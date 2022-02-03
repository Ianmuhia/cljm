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
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeWelcomeEmail, services.HandleVerifyEmailTask)
	mux.HandleFunc(TypeDailyVerse, services.HandleDailyVerseTask)
	//mux.HandleFunc(TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
