package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/4halavandala/l0/internal/entity"
	"github.com/4halavandala/l0/internal/handlers"
	"github.com/4halavandala/l0/pkg/cache"
	"github.com/4halavandala/l0/pkg/repository"
	"github.com/4halavandala/l0/pkg/server"
	"github.com/4halavandala/l0/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	// init config and env files
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config file %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error initializing .env file: %s", err.Error())
	}
	// db connection
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to init database: %s", err.Error())
	}
	// initializing dependencies
	repos := repository.NewRepository(db)
	cache := cache.NewCache()
	services := service.NewService(repos, cache)
	handlers := handlers.NewHandler(services)

	err = services.InitCache()
	if err != nil {
		logrus.Fatalf("Error initializing cache: %s", err.Error())
	}

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Info("Application started successfully")
	// nats connection
	sc, _ := stan.Connect(
		viper.GetString("nats.cluster_id"),
		viper.GetString("nats.client_id"),
		stan.NatsURL(viper.GetString("nats.url")))
	defer sc.Close()
	// receiving messages + saving data to db
	_, err = sc.Subscribe(viper.GetString("nats.queue"), func(msg *stan.Msg) {

		fmt.Println("Message received:\n" + string(msg.Data))

		var message entity.Model
		json.Unmarshal(msg.Data, &message)

		_, err := repos.Create(message)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		err = services.InitCache()
		if err != nil {
			logrus.Errorf("Error occured on saving data in cache: %s", err.Error())
		}
	}, stan.DurableName("my-durable"))
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
