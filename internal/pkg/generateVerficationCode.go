package pkg

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func GenerateCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}

func GenerateVerficationCode() int {
	//random genarator
	randomCode := GenerateCode()
	verficationCo := strconv.Itoa(randomCode)
	// Sender's email address and password
	from := "sangameshsnu@gmail.com"
	password := "snbk rmll zwdq uamx"

	// Recipient's email address
	to := "sangameshsasnur007@gmail.com"

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// Message content
	message := []byte("Subject: Verfication Code\n\n Verfication Code:" + verficationCo)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return -1
	}

	fmt.Println("Email sent successfully!")
	return randomCode
}
