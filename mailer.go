package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Email struct {
	Config
	To      []string
	Subject string
	Body    string
}

func New(config Config) *Email {
	return &Email{
		Config: Config{
			Host:     config.Host,
			Port:     config.Port,
			Username: config.Username,
			Password: config.Password,
			From:     config.From,
		},
	}
}

func Header(e *Email) map[string]string {
	header := make(map[string]string)
	header["From"] = e.Config.From
	header["To"] = strings.Join(e.To, ",")
	header["subject"] = e.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	return header
}

func (c *Email) Send() {
	msg := ""
	for key, val := range Header(c) {
		msg += fmt.Sprintf("%s: %s\r\n", key, val)
	}
	msg += "\r\n"
	msg += c.Body
	err := smtp.SendMail(c.Config.Host+":"+c.Config.Port,
		smtp.PlainAuth("", c.Config.Username, c.Config.Password, c.Config.Host),
		c.Config.From, c.To, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	log.Print("message sent")
}

func (e *Email) SetTo(to ...string) {
	e.To = to
}

func (e *Email) SetSubject(subject string) {
	e.Subject = subject
}

func (e *Email) SetBody(body string) {
	e.Body = body
}
