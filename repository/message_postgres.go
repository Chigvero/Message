package repository

import (
	"fmt"

	"github.com/Chigvero/Messageio/model/message"
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

func (r *MessagePostgresDB) CreateMessage(message modelMessage.Message) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (text) VALUES ($1) RETURNING id", messagesTable)

	if err := r.db.QueryRow(query, message.Text).Scan(&message.Id); err != nil {
		return 0, err
	}
	return message.Id, nil
}

func (r *MessagePostgresDB) ProcessMessage() {

}

func (r *MessagePostgresDB) GetMessageById(id int) (modelMessage.Message, error) {
	query := fmt.Sprintf("SELECT * from %s where id=$1", messagesTable)
	msg := modelMessage.Message{}
	err := r.db.QueryRow(query, id).Scan(&msg.Id, &msg.Text, &msg.CreatedAt, &msg.Processed)
	if err != nil {
		return modelMessage.Message{}, err
	}
	return msg, nil
}
