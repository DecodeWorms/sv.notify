package email

import (
	sendgridmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// client sender data
var (
	DefaultSenderEmailAddress = &sendgridmail.Email{Name: "Julia", Address: "amthetechguy@gmail.com"}
	defaultSenderName         = "DecodeWorms"
	defaultEmailAddress       = "amthetechguy@gmail.com"
)
