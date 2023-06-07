package utils

import (
	"doYourLogin/source/configuration"
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
		Host:     configuration.SMTP.ValueAsString(),
		Port:     configuration.PORT.ValueAsInt(),
		Username: configuration.EMAIL.ValueAsString(),
		Password: configuration.PASSWORD.ValueAsString(),
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
