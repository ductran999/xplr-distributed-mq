package main

import (
	"context"
	"log"
	"xplr-distributed-mq/evtstream"
	kafka "xplr-distributed-mq/evtstream/kafka/franzgo"
)

func main() {
	config := &kafka.Config{
		Brokers:                []string{"localhost:29092"},
		KafkaVersion:           "4.0.0.0",
		MaxRetry:               1,
		AllowAutoTopicCreation: true,
		EnableDebug:            true,
	}

	topic := "franz-go-topic"

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Close() //nolint

	msg := &evtstream.Message{
		Topic: topic,
		Key:   "user",
		Value: []byte("hello kafka from franzgo"),
	}

	if err := producer.Publish(context.Background(), msg); err != nil {
		log.Println("failed to send message:" + err.Error())
	} else {
		log.Println("message sent to topic:" + topic)
	}
}
