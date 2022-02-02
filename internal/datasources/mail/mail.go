package mail_client

import (
	"fmt"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var (
	MailServer *mail.SMTPServer
	MailClient *mail.SMTPClient
)

func GetMailServer() {
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
	MailClient = client
	MailServer = server

}
