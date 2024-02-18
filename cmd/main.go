package main

import (
	"log"
	"os"

	"github.com/OurLuv/l0/internal/pubsub"
	"github.com/OurLuv/l0/internal/storage/postgres"
	"github.com/nats-io/stan.go"
)

func main() {
	//* nats streaming
	sc, err := stan.Connect("my_cluster", "cl1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// creating and starting new subscriber
	s := pubsub.NewSubscriber(sc)
	s.Start()
	// creating and starting new subscriber
	p := pubsub.NewPublisher(sc)
	p.Start()

	//* storage
	dbPath, exists := os.LookupEnv("DB_PATH")
	if !exists {
		log.Printf("db path is not set: %s", dbPath)
	}
	pool, err := postgres.NewPool("")
	if err != nil {
		log.Printf("failed to init storage: %s", err)
		os.Exit(1)
	}
	defer pool.Close()
	log.Printf("Storage init")

}
