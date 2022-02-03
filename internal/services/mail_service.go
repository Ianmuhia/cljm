package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	mail_client "maranatha_web/internal/datasources/mail"
	redis_db "maranatha_web/internal/datasources/redis"
	tasks_client "maranatha_web/tasks/client"
	"os"

	"github.com/hibiken/asynq"
	mail "github.com/xhit/go-simple-mail/v2" //nolint:goimports
)

var (
	MailService mailServiceInterface = &mailService{}
)

// A list of task types.
const (
	TypeWelcomeEmail = "email:welcome"
	TypeDailyVerse   = "dailyVerse"
)

type mailService struct{}

type mailServiceInterface interface {
	SendMsg(m Mail) error
	VerifyMailCode(key string) string
	RemoveMailCode(key string)
}
type MailType int

const (
	MailConfirmation MailType = iota + 1
	PassReset
)

type MailData struct {
	Username string
	Code     string
}

type Mail struct {
	To      string
	From    string
	Subject string
	Content string
}

func (s *mailService) SendMsg(m Mail) error {
	log.Println("send message has been called.")
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}
	//return

	t1 := asynq.NewTask(TypeWelcomeEmail, marshal)
	info, err := tasks_client.TasksClient.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
	return nil
}

func HandleVerifyEmailTask(ctx context.Context, t *asynq.Task) error {
	var m Mail

	if err := json.Unmarshal(t.Payload(), &m); err != nil {
		return err
	}
	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject).SetBody(mail.TextPlain, m.Content)
	mc := mail_client.GetMailServer()
	err := email.Send(mc)

	if err != nil {
		log.Println("failing here.")
		log.Println(err)
		return err
	}
	log.Printf(" [*] Send Welcome Email to User %s", m.To)
	return nil
}

func (s *mailService) VerifyMailCode(key string) string {
	data, err := redis_db.RedisClient.Get(context.TODO(), key).Result()
	log.Printf("Redis Code %v or %v", data, err)
	if err != nil {
		log.Println(err)
		return "Invalid key provided or key not found"
	}
	return data
}

func (s *mailService) RemoveMailCode(key string) {

	if len(os.Args) > 1 {
		key = os.Args[1]
	}
	var foundedRecordCount int = 0
	iter := redis_db.RedisClient.Scan(context.Background(), 0, key, 0).Iterator()

	fmt.Printf("YOUR SEARCH PATTERN= %s\n", key)

	for iter.Next(context.Background()) {
		fmt.Printf("Deleted= %s\n", iter.Val())
		redis_db.RedisClient.Del(context.Background(), iter.Val())
		foundedRecordCount++
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Deleted Count %d\n", foundedRecordCount)
	err := redis_db.RedisClient.Close()
	if err != nil {
		log.Println(err)
		return
	}
}
