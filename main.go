package main

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient(apiPublicKeyEmail, apiPrivateKeyEmail)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "l6yaldbk6@relay.firefox.com",
				Name:  "Francis",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "l6yaldbk6@relay.firefox.com",
					Name:  "Francis",
				},
			},
			Subject:  "Greetings from Mailjet.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
