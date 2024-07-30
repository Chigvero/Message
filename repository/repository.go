package repository

import (
	"github.com/Chigvero/Messageio/modelMessage"
	"github.com/jmoiron/sqlx"
)

type MessageRepository interface {
	CreateMessage(message modelMessage.Message) (int, error)
	ProcessMessage()
	GetMessageById(id int) (modelMessage.Message, error)
}

type Repository struct {
	MessageRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MessageRepository: NewMessagePostgresDB(db),
	}
}
