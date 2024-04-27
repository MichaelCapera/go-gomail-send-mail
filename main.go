package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	// Configure the SMTP client to send email
	d := gomail.NewDialer("smtp.example.com", 587, "user@example.com", "xxxxPasswordxxxx")

	// Send e-mail
	err := sendEmail(d)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
	log.Println("Email sent successfully!")
}

func sendEmail(d *gomail.Dialer) error {
	// Create an email message
	m := gomail.NewMessage()
	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Hello from Go")
	m.SetBody("text/plain", "Hello, World!")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
