package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {

	sc, err := stan.Connect("test-cluster", "client-test", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	fileContent, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatal(err)
	}

	err = sc.Publish("order", fileContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data sended!")
}
