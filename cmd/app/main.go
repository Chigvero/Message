package main

import (
	"github.com/Chigvero/Messageio/internal/handler"
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
	DB, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(DB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	router := handlers.InitRoutes()
	http.ListenAndServe("localhost:8080", router)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
