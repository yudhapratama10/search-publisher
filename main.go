package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	kafkaClient "github.com/yudhapratama10/search-publisher/infrastructures/kafka"
	pgClient "github.com/yudhapratama10/search-publisher/infrastructures/pg"
)

func main() {

	dbClient, err := pgClient.InitClient()
	if err != nil {
		log.Fatal(err)
	}

	kClient := kafkaClient.InitClient()
	defer kClient.Close()

	err = kClient.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "test-messages",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Topic: "test-messages",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Topic: "test-messages",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := kClient.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
