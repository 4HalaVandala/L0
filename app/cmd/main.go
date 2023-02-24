package main

import (
	"github.com/4halavandala/l0/app/internal/handlers"
	"github.com/4halavandala/l0/app/pkg/repository"
	"github.com/4halavandala/l0/app/pkg/service"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config file %s", err)
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("Error initializing database %s", err)
	}

	_, err = stan.Connect("test-cluster", "client1")
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	if err := handlers.InitRoutes(); err != nil {
		logrus.Fatalf("Erorr initializing routes %s", err)
	}
	http.ListenAndServe(viper.GetString("port"), nil)
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
