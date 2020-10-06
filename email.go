package main

import (
	"bytes"
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"html/template"
	"log"
	"strconv"
)

const (
	senderEmail   = "l6yaldbk6@relay.firefox.com"  // the sender's email address can be loaded from a config file or fetched from db
	senderName    = "Francis"                      //  name can also be loaded from config or fetched from a DB
	templateFPath = "./assets/email-template.html" //  path to the html template
)

// sendEmail formats an email sent to recipientEmail to inform them that an alert is raised at time ts, indicating the id of the alert.
func sendEmail(recipientEmail, recipientName, ts string, id int, imgPath string) {
	mailjetClient := mailjet.NewMailjetClient(apiPublicKeyEmail, apiPrivateKeyEmail)
	header := fmt.Sprintf("Someone is not wearing their mask - alert #%d", id)
	body := generateBody(templateFPath, recipientName, ts, id, imgPath)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: senderEmail,
				Name:  senderName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: recipientEmail,
					Name:  recipientName,
				},
			},
			Subject:  header,
			TextPart: header,
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

func generateBody(templateEmailFileName, recipientName, ts string, alertId int, imgPath string) string {
	var data struct {
		RecipientName template.HTML
		EventTime     template.HTML
		AlertId       string
		PathToImage   string
	}
	data.RecipientName = template.HTML(recipientName)
	data.EventTime = template.HTML(ts)
	data.AlertId = strconv.Itoa(alertId)
	data.PathToImage = imgPath

	t, err := template.ParseFiles(templateEmailFileName)
	if err != nil {
		log.Fatalln(err)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}
