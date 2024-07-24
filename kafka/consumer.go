package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func NewConsumer() (*kafka.ConsumerGroup, error) {
	topic := "message"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp",
		"localhost:9092", topic, partition)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return nil, err
}
