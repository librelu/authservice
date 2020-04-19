package smtp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"text/template"
)

// NewSMTPClient create a new SMTP client
func NewSMTPClient(identity, username, password, host, port string) Handler {
	auth := smtp.PlainAuth(identity, username, password, host)
	return &client{
		smtpAuth: auth,
		host:     host,
		port:     port,
		username: username,
	}
}

// SendEmail send email from client
func (c *client) SendEmail(emails []string, message []byte) error {
	err := smtp.SendMail(
		c.host+":"+c.port,
		c.smtpAuth,
		c.username,
		emails,
		message)
	if err != nil {
		return err
	}
	return nil
}

// GetWelcomeEmail get email email template
func (c *client) SendWelcomeEmail(email, username, coupon string) error {
	f, err := os.Open(fmt.Sprintf(templatePath, welcomeFile))
	if err != nil {
		return err
	}
	welcomeTemp, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	template, err := template.New("welcome_letter").Option("missingkey=zero").Parse(string(welcomeTemp))
	if err != nil {
		return err
	}

	funcMap := make(map[string]interface{})
	funcMap["Username"] = username
	funcMap["Coupon"] = coupon
	funcMap["Subject"] = welcomeTitle
	funcMap["Email"] = email

	message := new(bytes.Buffer)
	err = template.Execute(message, funcMap)
	if err != nil {
		return err
	}

	err = c.SendEmail([]string{email}, message.Bytes())
	if err != nil {
		return err
	}
	return nil
}
