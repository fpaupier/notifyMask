package main

import (
	"bytes"
	"encoding/binary"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

const (
	topic      = "to-notify-topic"             // To consume from
	adminEmail = "dluumi0ke@relay.firefox.com" // Email address to which send the email - could be loaded from config or a DB
	adminName  = "Camille"                     // Name of the admin
	imgPath    = "./assets/img.jpeg"           // image responsible from the alert
)

func main() {
	// Consume queue of alerts
	consumer := getConsumer(topic)
	defer consumer.Close()

	for {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			alertId := getAlertId(e.Value)
			// Retrieve the record
			ts := getAlertEventTime(alertId)
			fetchImage(alertId, imgPath)
			// Prepare email for admin
			// Send email to admin
			sendEmail(adminEmail, adminName, ts, alertId, imgPath)
			// Update status of alert to sent
			checkAlert(alertId)
		case kafka.PartitionEOF:
			log.Printf("%% Reached %v\n", e)
		case kafka.Error:
			log.Fatalf("%% Error: %v\n", e)
		}
	}
}

// getAlertId decode the binary representation of the alert id used over the wire into an integer.
func getAlertId(value []byte) int {
	var r = bytes.NewReader(value)
	var id uint64
	if err := binary.Read(r, binary.LittleEndian, &id); err != nil {
		log.Fatalf("failed to decode alert id: %v\n", err)
	}
	return int(id)
}
