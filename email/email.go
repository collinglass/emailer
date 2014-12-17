package email

import (
	"fmt"
	"net/smtp"
	"strconv"
)

// Email Client
// Currently Setup as GMAIL
type Emailer struct {
	username    string
	password    string
	emailServer string
	port        int
}

func Setup(email, password string) *Emailer {
	return &Emailer{
		email,
		password,
		"smtp.gmail.com",
		587,
	}
}

func (e *Emailer) SendMail(to string, subject, message string) error {
	var err error

	auth := smtp.PlainAuth("",
		e.username,
		e.password,
		e.emailServer,
	)

	err = smtp.SendMail(
		e.emailServer+":"+strconv.Itoa(e.port), // in our case, "smtp.google.com:587"
		auth,
		e.username,
		[]string{to},
		[]byte(fmt.Sprintf("Subject: %s\r\n\r\n", subject)+message))

	return err
}
