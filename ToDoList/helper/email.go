package helper

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendGomail(templatePath string, userEmail, userName, taskTitle, taskDescription, dueDate, timeRemaining string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	} 

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	
	err = t.Execute(&body, struct {
		UserName      string
		TaskTitle     string
		TaskDescription string
		DueDate       string
		TimeRemaining string
	}{
		UserName:      userName,
		TaskTitle:     taskTitle,
		TaskDescription: taskDescription,
		DueDate:       dueDate,
		TimeRemaining: timeRemaining,
	})
	if err != nil {
		panic(err)
	}

	email := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	port,err := strconv.Atoi(smtpPort)
	PanicIfError(err)

	m := gomail.NewMessage()
	m.SetHeader("From", email) 
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Task Reminder")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(smtpHost, port, email, password) 
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
