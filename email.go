package main

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
)

const (
	senderEmail = "l6yaldbk6@relay.firefox.com" // the sender's email address can be loaded from a config file or fetched from db
	senderName  = "Francis"                     //  name can also be loaded from config or fetched from a DB
)

// sendEmail formats an email sent to recipientEmail to inform them that an alert is raised at time ts, indicating the id of the alert.
func sendEmail(recipientEmail, ts string, id int) {
	mailjetClient := mailjet.NewMailjetClient(apiPublicKeyEmail, apiPrivateKeyEmail)
	body := fmt.Sprintf("<h3>Alert</h3><br />Someone has been seen not wearing their mask at %s. Check alert ID#%d for more information.", ts, id)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: senderEmail,
				Name:  senderName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: recipientEmail,
					Name:  senderName,
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
