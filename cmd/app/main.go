package main

import (
	"github.com/Chigvero/Messageio/internal/handler"
	"github.com/Chigvero/Messageio/internal/kafka"
	"github.com/Chigvero/Messageio/internal/repository"
	"github.com/Chigvero/Messageio/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	err := InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	config := &repository.Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetString("db.ssl_mode"),
	}
	conn, err := repository.NewPostgresConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(conn)
	producer, err := kafka.NewProducer()
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := kafka.NewConsumerGroup(repos)
	if err != nil {
		log.Fatal(err)
	}
	go consumer.Start()
	services := service.NewService(repos, producer)
	handlers := handler.NewHandler(services)
	router := handlers.InitRoutes()
	go func() {
		if err := http.ListenAndServe("localhost:8080", router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()
	consumer.Close()
	for {

	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
