package kafka

import (
	"encoding/json"
	"fmt"
	Message "github.com/Chigvero/Messageio"
	"github.com/Chigvero/Messageio/internal/repository"
	"github.com/IBM/sarama"
	"log"
)

type Consumer struct {
	consumerGroup sarama.Consumer
	repos         repository.MessageRepository
}

func NewConsumerGroup(repos *repository.Repository) (*Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		consumer,
		repos.MessageRepository,
	}, nil
}

func (c *Consumer) Start() error {
	partConsumer, err := c.consumerGroup.ConsumePartition("my-topic", 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	log.Println("Consumer STARTED")
	defer partConsumer.Close()
	for {
		select {
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				return fmt.Errorf("message channel closed")
			}
			var message Message.Message
			//fmt.Println(string(msg.Value))
			if err := json.Unmarshal(msg.Value, &message); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			c.repos.ProcessMessage(message.Id)
			//log.Printf("Processed message: %v\n", message)
		}
	}
}

func (c *Consumer) Close() {
	c.consumerGroup.Close()
}
