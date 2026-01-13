package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:29092"}
	topic := "test-topic"

	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// debug
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("user-1"),
		Value: sarama.StringEncoder("hello kafka from sarama"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	log.Printf("message sent to partition %d at offset %d\n", partition, offset)
}
