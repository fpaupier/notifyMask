package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

const (
	bootstrapServers = ConfluentServer
	ccloudAPIKey     = ConfluentApiKey
	ccloudAPISecret  = ConfluentSecret
)

// getConsumer returns a Kafka consumer for the given topic.
func getConsumer(topic string) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"sasl.mechanisms":   "PLAIN",
		"security.protocol": "SASL_SSL",
		"sasl.username":     ccloudAPIKey,
		"sasl.password":     ccloudAPISecret,
		"group.id":          "alert-consumer",
		"auto.offset.reset": "earliest"})
	if err != nil {
		log.Fatalf("Failed to connect to Kafka %v", err)
	}

	// Read from Kafka
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("failed to subscribe to topic: %v\n", err)
	}
	return consumer
}
