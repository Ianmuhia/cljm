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

	//	t1, err := task.NewWelcomeEmailTask(42)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	t2, err := task.NewReminderEmailTask(42)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// Process the task immediately.
	//	info, err := client.Enqueue(t1)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Printf(" [*] Successfully enqueued task: %+v", info)
	//
	//	// Process the task 24 hours later.
	//	info, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Printf(" [*] Successfully enqueued task: %+v", info)
	//}
}

func GetTasksClient() *asynq.Client {
	return TasksClient
}
