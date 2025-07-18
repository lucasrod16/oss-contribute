package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type email struct {
	from     string
	to       string
	smtpHost string
	smtpPort string
	subject  string
	body     string
}

func (e email) Send(password string) error {
	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"From: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n",
		e.to, e.from, e.subject, e.body,
	))
	auth := smtp.PlainAuth("", e.from, password, e.smtpHost)
	return smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.from, []string{e.to}, msg)
}

func main() {
	password := os.Getenv("GMAIL_APP_PASSWORD")
	workflowURL := os.Getenv("WORKFLOW_URL")
	if password == "" {
		log.Fatal("GMAIL_APP_PASSWORD environment variable must be set.")
	}
	if workflowURL == "" {
		log.Fatal("WORKFLOW_URL environment variable must be set.")
	}
	email := email{
		from:     "lucas.rodriguez9616@gmail.com",
		to:       "lucas.rodriguez9616@gmail.com",
		smtpHost: "smtp.gmail.com",
		smtpPort: "587",
		subject:  "OSS Projects GitHub Workflow Failure",
		body: fmt.Sprintf(
			"The oss-projects cronjob workflow did not succeed. "+
				"The GitHub data in the lucasrod16-github-data GCS bucket was not updated.\n\n"+
				"Workflow URL: %s",
			workflowURL,
		),
	}
	if err := email.Send(password); err != nil {
		log.Fatalf("failed to send email: %v", err)
	}
	log.Println("Email sent successfully!")
}
