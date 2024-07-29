package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Chigvero/Messageio/internal/repository"
	Message "github.com/Chigvero/Messageio/modelMessage"
	"github.com/IBM/sarama"
)

type Consumer struct {
	consumerGroup sarama.Consumer
	repos         repository.MessageRepository
}

func NewConsumerGroup(repos *repository.Repository) (*Consumer, error) {
	fmt.Println("kafka:9092-consumer")
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
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
		fmt.Println("err: %s", err.Error())
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
