package kafka

import (
	"encoding/json"
	"fmt"

	Message "github.com/Chigvero/Messageio/model/message"
	"github.com/IBM/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer() (*Producer, error) {
	fmt.Println("kafka:9092-producer")
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		return nil, err
	}
	return &Producer{
		syncProducer: producer,
	}, nil
}

func (p *Producer) SendMessage(message Message.Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: "my-topic",
		Value: sarama.ByteEncoder(bytes),
	}
	_, _, err = p.syncProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) Close() {
	p.syncProducer.Close()
}
