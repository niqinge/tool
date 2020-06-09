package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type IEmailHelper interface {
	Send(option *Option) error
}

type EmailHelper struct {
	*Config
}

func NewEmailHelper(config *Config) *EmailHelper {
	return &EmailHelper{Config: config}
}

func (h *EmailHelper) Send(option *Option) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s<%s>", h.UserName, h.From))
	m.SetHeader("To", option.To...)
	m.SetHeader("Cc", option.Cc...)
	m.SetHeader("Subject", option.Subject)
	m.SetBody("text/plain", option.Content)

	d := gomail.NewDialer(h.Host, h.Post, h.From, h.Password)
	return d.DialAndSend(m)
}

