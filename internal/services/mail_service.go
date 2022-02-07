package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hibiken/asynq"
	mail "github.com/xhit/go-simple-mail/v2" //nolint:goimports

	mail_client "maranatha_web/internal/datasources/mail"
	redis_db "maranatha_web/internal/datasources/redis"
	tasks_client "maranatha_web/tasks/client"
)

// TypeWelcomeEmail A list of task types.
const (
	TypeWelcomeEmail     = "email:welcome"
	VerificationDataType = ""
	//TypeDailyVerse   = "dailyVerse"
)

type mailService struct{}

type MailService interface {
	SendMsg(m *Mail) error
	VerifyMailCode(key string) string
	RemoveMailCode(key string)
	NewMail(from string, to string, subject string, mailType MailType, data *MailData) *Mail
}

func NewMailService() MailService {
	return &mailService{}
}

type MailType int

const (
	MailConfirmation MailType = iota + 1
	PassReset
)

// MailData represents the data to be sent to the template of the mail.
type MailData struct {
	Username string
	Code     string
}

// Mail represents a email request
type Mail struct {
	from     string
	to       string
	subject  string
	body     string
	mailType MailType
	data     *MailData
}

// VerificationData represents the type for the data stored for verification.
type VerificationData struct {
	Email     string    `json:"email" validate:"required" sql:"email"`
	Code      string    `json:"code" validate:"required" sql:"code"`
	ExpiresAt time.Time `json:"expires_at" sql:"expiresat"`
	Type      string    `json:"type" sql:"type"`
}

func (ms *mailService) SendMsg(m *Mail) error {

	log.Println("send message has been called.")

	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}

	t1 := asynq.NewTask(TypeWelcomeEmail, marshal)
	info, err := tasks_client.TasksClient.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
	return nil
}

// NewMail returns a new mail request.
func (ms *mailService) NewMail(from string, to string, subject string, mailType MailType, data *MailData) *Mail {
	return &Mail{
		from:     from,
		to:       to,
		subject:  subject,
		mailType: mailType,
		data:     data,
	}
}

func HandleVerifyEmailTask(ctx context.Context, t *asynq.Task) error {
	var m Mail

	if err := json.Unmarshal(t.Payload(), &m); err != nil {
		return err
	}
	log.Println(m)
	email := mail.NewMSG()
	email.SetFrom(m.from).AddTo(m.to).SetSubject(m.subject).SetBody(mail.TextPlain, m.body)
	mc := mail_client.GetMailServer()
	err := email.Send(mc)

	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf(" [*] Send Welcome Email to User %s", m.to)
	return nil
}

func (ms *mailService) VerifyMailCode(key string) string {
	data, err := redis_db.RedisClient.Get(context.TODO(), key).Result()
	log.Printf("Redis Code %v or %v", data, err)
	if err != nil {
		log.Println(err)
		return "Invalid key provided or key not found"
	}
	return data
}

func (ms *mailService) RemoveMailCode(key string) {

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
