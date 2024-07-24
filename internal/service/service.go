package service

import (
	Intern "github.com/Chigvero/Messageio"
	repository "github.com/Chigvero/Messageio/internal/repository"
)

type Business interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage()
	GetMessageById(id int) (Intern.Message, error)
}

type Service struct {
	Business
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repos:    repos,
		Business: NewMessage(repos),
	}
}
