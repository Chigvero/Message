package service

import (
	Intern "github.com/Chigvero/Messageio"
	"github.com/Chigvero/Messageio/internal/repository"
)

type MessageService struct {
	repos repository.MessageRepository
}

func NewMessageService(repos *repository.Repository) *MessageService {
	return &MessageService{
		repos: repos.MessageRepository,
	}
}

func (s *MessageService) CreateMessage(message Intern.Message) (int, error) {
	return s.repos.CreateMessage(message)
}

func (s *MessageService) ProcessMessage() {
	//return s.repos.ProcessMessage()
}

func (s *MessageService) GetMessageById(id int) (Intern.Message, error) {
	return s.repos.GetMessageById(id)
}
