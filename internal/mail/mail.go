package mail

import (
	"bytes"
	"net/smtp"
	"os"
)

func Send(subject string, body bytes.Buffer, to []string) error {
	// Authorise Sender Account (currently using personal account)
	auth := smtp.PlainAuth(
		"",
		"kyegregory001@gmail.com",
		os.Getenv("MAIL_APP_PASSWORD"),
		"smtp.gmail.com",
	)

	// Define Meta-Data
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	// Send Main
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"kyegregory001@gmail.com",
		to,
		[]byte(msg),
	)

	// Check Errors
	if (err != nil) { return err}
	return nil
}