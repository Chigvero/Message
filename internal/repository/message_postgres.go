package repository

import (
	"fmt"
	Intern "github.com/Chigvero/Messageio"
	"github.com/jmoiron/sqlx"
)

type MessagePostgresDB struct {
	db *sqlx.DB
}

func NewMessagePostgresDB(db *sqlx.DB) *MessagePostgresDB {
	return &MessagePostgresDB{
		db: db,
	}
}

func (r *MessagePostgresDB) CreateMessage(message Intern.Message) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (text) VALUES ($1) RETURNING id", messagesTable)

	if err := r.db.QueryRow(query, message.Text).Scan(&message.Id); err != nil {
		return 0, err
	}
	return message.Id, nil
}

func (r *MessagePostgresDB) ProcessMessage(id int) {
	query := fmt.Sprintf("UPDATE %x SET processed=true where id=$1", messagesTable)
	r.db.Exec(query, id)
}

func (r *MessagePostgresDB) GetMessageById(id int) (Intern.Message, error) {
	query := fmt.Sprintf("SELECT * from %s where id=$1", messagesTable)
	msg := Intern.Message{}
	err := r.db.QueryRow(query, id).Scan(&msg.Id, &msg.Text, &msg.CreatedAt, &msg.Processed)
	if err != nil {
		return Intern.Message{}, err
	}
	return msg, nil
}
