package Message

import "time"

type Message struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Processed bool      `json:"processed"`
}
