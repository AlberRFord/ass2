package utils

import (
	"fmt"
	"net/smtp"
)

// SendConfirmationEmail sends a confirmation email to the user
func SendConfirmationEmail(email, token string) error {
	from := "your.email@example.com"
	password := "your-email-password"
	to := email

	subject := "Confirm Your Registration"
	body := fmt.Sprintf("Please click the following link to confirm your registration: http://yourwebsite.com/confirm?token=%s", token)

	auth := smtp.PlainAuth("", from, password, "smtp.example.com")

	err := smtp.SendMail("smtp.example.com:587", auth, from, []string{to}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
