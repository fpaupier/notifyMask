package main

import (
	"bytes"
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"html/template"
	"log"
	"strconv"
	"time"
)

const (
	senderEmail   = "l6yaldbk6@relay.firefox.com"  // the sender's email address can be loaded from a config file or fetched from db
	senderName    = "Francis"                      //  name can also be loaded from config or fetched from a DB
	templateFPath = "./assets/email-template.html" //  path to the html template
)

// sendEmail formats an email sent to recipientEmail to inform them that an alert is raised at time ts, indicating the id of the alert.
func sendEmail(recipientEmail, recipientName, ts string, id int) {
	mailjetClient := mailjet.NewMailjetClient(apiPublicKeyEmail, apiPrivateKeyEmail)
	header := fmt.Sprintf("Someone's not wearing its mask - Alert #%d", id)
	body := generateBody(templateFPath, recipientName, ts, id)
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
	log.Printf("mail sent for alert id #%d: %+v\n", id, res)
}

func generateBody(templateEmailFileName, recipientName, ts string, alertId int) string {
	var data struct {
		RecipientName template.HTML
		EventTime     template.HTML
		AlertId       string
	}
	data.RecipientName = template.HTML(recipientName)
	data.AlertId = strconv.Itoa(alertId)
	ti, err := time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		log.Fatalf("failed to parse time: %v", err)
	}
	data.EventTime = template.HTML(ti.Format(time.RFC850))

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
