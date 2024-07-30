package service

import (
	"github.com/Chigvero/Messageio/internal/repository"
	"github.com/Chigvero/Messageio/modelMessage"
)

type MessageService struct {
	repos repository.MessageRepository
}

func NewMessageService(repos *repository.Repository) *MessageService {
	return &MessageService{
		repos: repos.MessageRepository,
	}
}

func (s *MessageService) CreateMessage(message modelMessage.Message) (int, error) {
	return s.repos.CreateMessage(message)
}

func (s *MessageService) ProcessMessage() {
	//return s.repos.ProcessMessage()
}

func (s *MessageService) GetMessageById(id int) (modelMessage.Message, error) {
	return s.repos.GetMessageById(id)
}
