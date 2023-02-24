package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	sc, err := stan.Connect("test-cluster", "client1")
	if err != nil {
		log.Fatal(err)
	}
	_, err = sc.Subscribe("foo", func(msg *stan.Msg) {
		// Handle message here
		fmt.Println(string(msg.Data))
	}, stan.DurableName("my-durable"))
	if err != nil {
		log.Fatal(err)
	}
}
