package service

import (
	Intern "github.com/Chigvero/Messageio"
	"github.com/Chigvero/Messageio/internal/repository"
)

type Message struct {
	repos *repository.Repository
}

func NewMessage(repos *repository.Repository) *Message {
	return &Message{
		repos: repos,
	}
}

func (s *Message) CreateMessage(message Intern.Message) (int, error) {
	return s.repos.CreateMessage(message)
}

func (s *Message) ProcessMessage() {
	//return s.repos.ProcessMessage()
}

func (s *Message) GetMessageById(id int) (Intern.Message, error) {
	return s.repos.GetMessageById(id)
}
