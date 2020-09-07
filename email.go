package main

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
)

func sendEmail(to, ts string, id int) {
	mailjetClient := mailjet.NewMailjetClient(apiPublicKeyEmail, apiPrivateKeyEmail)
	body := fmt.Sprintf("<h3>Alert</h3><br />Someone has been seen not wearing their mask at %s. Check alert ID#%d for more information.", ts, id)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "l6yaldbk6@relay.firefox.com",
				Name:  "Francis",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  "Francis",
				},
			},
			Subject:  "Mask Alert",
			TextPart: "Someone is not wearing their mask",
			HTMLPart: body,
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data: %+v\n", res)
}
