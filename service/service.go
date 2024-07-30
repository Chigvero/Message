package service

import (
	"github.com/Chigvero/Messageio/internal/repository"
	"github.com/Chigvero/Messageio/modelMessage"
)

type Message interface {
	CreateMessage(message modelMessage.Message) (int, error)
	ProcessMessage()
	GetMessageById(id int) (modelMessage.Message, error)
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repos),
	}
}
