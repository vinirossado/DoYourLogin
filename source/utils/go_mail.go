package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

func InitEmailServer() *EmailSender {
	sender := &EmailSender{
		Host:     "smtp.gmail.com",
		Port:     465,
		Username: "vinirossado@gmail.com",
		Password: "anpch@example.com",
	}

	return sender
}

func (es *EmailSender) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", es.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(es.Host, es.Port, es.Username, es.Password)

	err := d.DialAndSend(m)

	if err != nil {
		return fmt.Errorf("Error sending email: %v", err)
	}

	return nil
}
