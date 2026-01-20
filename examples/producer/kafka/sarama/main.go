package main

import (
	"context"
	"log"
	"xplr-distributed-mq/mq"
	kafka "xplr-distributed-mq/mq/kafka/sarama"
)

func main() {
	config := kafka.Config{
		Brokers:      []string{"localhost:29092"},
		KafkaVersion: "4.0.0.0",
		MaxRetry:     1,
		EnableDebug:  true,
	}

	topic := "test-topic"

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Close()

	msg := &mq.Message{
		Topic: topic,
		Key:   "user-1",
		Value: []byte("hello kafka from sarama"),
	}

	partition, offset, err := producer.Publish(context.Background(), msg)
	if err != nil {
		log.Printf("failed to send message: %v", err)
	} else {
		log.Printf("message sent to partition %d at offset %d\n", partition, offset)
	}
}
