package main

import (
	"log"
	"net/http"

	"github.com/Chigvero/Messageio/internal/handler"
	"github.com/Chigvero/Messageio/internal/kafka"
	"github.com/Chigvero/Messageio/internal/repository"
	"github.com/Chigvero/Messageio/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
	go func() {
		if err := consumer.Start(); err != nil {
			log.Fatalf("Failed to start consumer: %v", err)
		}
	}()
	services := service.NewService(repos, producer)
	handlers := handler.NewHandler(services)
	router := handlers.InitRoutes()
	go func() {
		if err := http.ListenAndServe(":8081", router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	// Ожидание сигнала завершения работы
	for {

	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
