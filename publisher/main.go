package main

import (
	"encoding/json"
	"fmt"
	"github.com/4halavandala/l0/internal/entity"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	sc, err := stan.Connect(viper.GetString("nats.cluster_id"), viper.GetString("nats.producer_id"), stan.NatsURL(viper.GetString("nats.url")))
	if err != nil {
		log.Fatalf("error connecting nats %s", err.Error())
	}
	defer sc.Close()

	jsonFile, err := os.Open(viper.GetString("json.path"))
	if err != nil {
		logrus.Fatalf("error openning json file: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var order entity.Model
	json.Unmarshal(byteValue, &order)
	fmt.Println(order)

	sc.Publish(viper.GetString("nats.queue"), byteValue)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
