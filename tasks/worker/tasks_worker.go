package main

import (
	"log"

	"github.com/hibiken/asynq" //nolint:goimports
	"maranatha_web/services"
	task "maranatha_web/tasks"
)

var (
	TasksServer *asynq.Server
)

// workers.go
func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	TasksServer = srv

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, services.HandleVerifyEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

//func GetTasksWorker() *asynq.Server {
//	return TasksServer
//}
