package mail

import (
	"bytes"
	"net/smtp"
)

func Send(subject string, body bytes.Buffer, to []string) error {
	auth := smtp.PlainAuth(
		"",
		"kyegregory001@gmail.com",
		"yazlazxbrbfzcxfg", // gmail app password
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"kyegregory001@gmail.com",
		to,
		[]byte(msg),
	)

	if (err != nil) { return err}
	return nil
}