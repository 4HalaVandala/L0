//package main
//
//import (
//	"fmt"
//	"github.com/nats-io/stan.go"
//	"log"
//)
//
//func main() {
//	sc, err := stan.Connect("test-cluster", "client1")
//	if err != nil {
//		log.Fatal(err)
//	}
//	_, err = sc.Subscribe("foo", func(msg *stan.Msg) {
//		// Handle message here
//		fmt.Println(string(msg.Data))
//	}, stan.DurableName("my-durable"))
//	if err != nil {
//		log.Fatal(err)
//	}
//}

package main

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config file %s", err)
	}
	_, err := stan.Connect(viper.GetString("nats.cluster_id"), viper.GetString("nats.client_id"))
	if err != nil {
		logrus.Fatalf("Error publishing message: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../L0/app/configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
