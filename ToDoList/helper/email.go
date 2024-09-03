package helper

import (
	"bytes"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendGomail(templatePath string, userEmail, userName, taskTitle, taskDescription, dueDate, timeRemaining string) {
	// Get the HTML template
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	// Execute template with data
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

	// Create and send email
	m := gomail.NewMessage()
	m.SetHeader("From", "test@gmail.com") // input ur email here
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Task Reminder")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "test@gmail.com", "mcanca") // input password here
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
