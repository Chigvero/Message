package kafka

import (
	"encoding/json"
	Message "github.com/Chigvero/Messageio"
	"github.com/Chigvero/Messageio/internal/repository"
	"github.com/IBM/sarama"
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
	message := Message.Message{}
	partConsumer, err := c.consumerGroup.ConsumePartition("my-topic", 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	defer partConsumer.Close()
	for {
		select {
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				return err
			}
			json.Unmarshal(msg.Value, &message)
			c.repos.ProcessMessage(message.Id)
		}
	}
}

func (c *Consumer) Close() {
	c.consumerGroup.Close()
}
