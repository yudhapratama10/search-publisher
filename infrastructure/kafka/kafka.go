package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func CheckConnection() (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", "localhost:9092")
}

func InitClient() *kafka.Writer {
	kafkaClient := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Balancer: &kafka.LeastBytes{},
	}

	return kafkaClient
}
