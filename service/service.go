package service

import (
	"github.com/Chigvero/Messageio/internal/repository"
	message "github.com/Chigvero/Messageio/model/message"
)

type Message interface {
	CreateMessage(message message.Message) (int, error)
	ProcessMessage()
	GetMessageById(id int) (message.Message, error)
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repos),
	}
}
