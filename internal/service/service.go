package service

import (
	Intern "github.com/Chigvero/Messageio"
	"github.com/Chigvero/Messageio/internal/repository"
)

type Message interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage()
	GetMessageById(id int) (Intern.Message, error)
}

type Service struct {
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Message: NewMessageService(repos),
	}
}
