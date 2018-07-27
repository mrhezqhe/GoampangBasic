package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go"
	"log"
)

func main() {

	from := mail.NewEmail("M Rizki", "testfrom@gmail.com")
	subject := "Golang with sendgrid"
	to := mail.NewEmail("Jhon Doe", "testto@gmail.com")
	plainTextContent := "Hi there, there are many things to do with Golang"
	htmlContent := "<strong>Happy Programming</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("THISISMYKEY")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}