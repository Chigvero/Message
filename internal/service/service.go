package service

import (
	"github.com/Chigvero/Messageio/internal/kafka"
	"github.com/Chigvero/Messageio/internal/repository"
	Intern "github.com/Chigvero/Messageio/modelMessage"
)

type Message interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage(id int)
	GetMessageById(id int) (Intern.Message, error)
	GetStats() (int, error)
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository, producer *kafka.Producer) *Service {
	return &Service{
		Message: NewMessageService(repos, producer),
	}
}
