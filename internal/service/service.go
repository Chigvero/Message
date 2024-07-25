package service

import (
	Intern "github.com/Chigvero/Messageio"
	"github.com/Chigvero/Messageio/internal/kafka"
	"github.com/Chigvero/Messageio/internal/repository"
)

type Message interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage(id int)
	GetMessageById(id int) (Intern.Message, error)
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository, producer *kafka.Producer) *Service {
	return &Service{
		Message: NewMessageService(repos, producer),
	}
}
