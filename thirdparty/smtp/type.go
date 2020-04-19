package smtp

import (
	"net/smtp"
	"os"
)

const (
	welcomeFile  = "welcome.tmpl"
	welcomeTitle = "Welcome to Amazing Talker"
)

var (
	templatePath = os.Getenv("GOPATH") + "/src/github.com/authsvc/thirdparty/smtp/template/%s"
)

type client struct {
	smtpAuth  smtp.Auth
	host      string
	port      string
	username  string
	fromEmail string
}

// Handler SMTP client handler
type Handler interface {
	SendEmail(emails []string, message []byte) error
	SendWelcomeEmail(email, username, coupon string) error
}
