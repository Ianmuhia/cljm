package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	mail "github.com/xhit/go-simple-mail/v2" //nolint:goimports
	redis_db "maranatha_web/datasources/redis"
)

var (
	MailService mailServiceInterface = &mailService{}
)

type mailService struct{}

type mailServiceInterface interface {
	SendMsg(m Mail) error
	ListenForMail()
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

var mailChan chan Mail

func (s *mailService) ListenForMail() {
	go func() {
		for {
			msg := <-mailChan
			err := s.SendMsg(msg)
			if err != nil {
				panic(err)
			}
		}
	}()
}

func (s *mailService) SendMsg(m Mail) error {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		fmt.Println(err)
	}

	email := mail.NewMSG()

	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject).SetBody(mail.TextPlain, m.Content)

	err = email.Send(client)

	if err != nil {
		log.Println(err)
		return err
	}
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
