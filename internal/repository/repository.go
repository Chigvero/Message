package repository

import (
	Intern "github.com/Chigvero/Messageio"
	"github.com/jmoiron/sqlx"
)

type DataBase interface {
	CreateMessage(message Intern.Message) (int, error)
	ProcessMessage()
	GetMessageById(id int) (Intern.Message, error)
}

type Repository struct {
	DataBase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DataBase: NewMessagePostgresDB(db),
	}
}
