package email

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"sv.notify/config"
	"sv.notify/model"
)

func SendWelcomeMessage(data model.Welcome) error {
	from := mail.NewEmail(defaultSenderName, defaultEmailAddress)
	to := mail.NewEmail(data.FirstName, data.Email)

	subject := "Welcome to the Barcelona Football Club"
	plainTextContent := "Welcome To the Barcelona Football Club"
	htmlContent := "<strong>Welcome To the Barcelona Football Clu</strong>"

	return processTheMail(from, to, subject, plainTextContent, htmlContent)
}

func processTheMail(from, to *mail.Email, subj, plainTextContent, htmlContent string) error {
	message := mail.NewSingleEmail(from, subj, to, plainTextContent, htmlContent)
	c := config.ImportConfig(config.Config{})
	client := sendgrid.NewSendClient(c.SendgridApiKey)
	client.Method = "POST"
	response, err := client.Send(message)
	if err != nil {
		return err
	} else {
		fmt.Println(response.StatusCode)
	}
	return nil
}
