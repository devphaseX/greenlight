package mailer

import (
	"bytes"
	"embed"
	"html/template"
	"time"

	"github.com/go-mail/mail/v2"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer struct {
	dailer *mail.Dialer
	sender string
}

func New(host string, port int, username, password, sender string) *Mailer {
	dailer := mail.NewDialer(host, port, username, password)
	dailer.Timeout = 5 * time.Second

	return &Mailer{
		dailer: dailer,
		sender: sender,
	}
}

func (m *Mailer) Send(recipient, templateFile string, data any) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)

	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(subject, "subject", data)

	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)

	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)

	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	for i := 1; i <= 3; i++ {
		err = m.dailer.DialAndSend(msg)
		if err == nil {
			return err
		}

		// If it didn't work, sleep for a short time and retry.
		time.Sleep(time.Millisecond * 500)
	}

	return nil
}
