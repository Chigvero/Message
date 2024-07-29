package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	messagesTable = "messages"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"ssl_mode"`
}

func NewPostgresConnection(cfg *Config) (*sqlx.DB, error) {
	query := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode)
	fmt.Println(query)
	db, err := sqlx.Open("postgres", query)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(1)
		return nil, err
	}
	return db, nil
}
