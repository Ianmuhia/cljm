package tasks_client

import (
	"github.com/hibiken/asynq"
)

var (
	TasksClient *asynq.Client
)

// client.go
func init() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	TasksClient = client

}

func GetTasksClient() *asynq.Client {
	return TasksClient
}
