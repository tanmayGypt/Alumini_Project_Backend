package helper

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(recipient string, subject string, body string) error {
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	// Convert SMTP_PORT to int
	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		return err
	}
	m := gomail.NewMessage()

	m.SetHeader("From", emailUser)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(smtpServer, smtpPortInt, emailUser, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
