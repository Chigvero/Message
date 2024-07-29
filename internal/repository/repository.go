package repository

import (
	Intern "github.com/Chigvero/Messageio/modelMessage"
	"github.com/jmoiron/sqlx"
)

type MessageRepository interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage(id int)
	GetMessageById(id int) (Intern.Message, error)
	GetStats() (int, error)
}

type Repository struct {
	MessageRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MessageRepository: NewMessagePostgresDB(db),
	}
}
