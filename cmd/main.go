package main

import (
	"log"

	"github.com/OurLuv/l0/internal/pubsub"
	"github.com/nats-io/stan.go"
)

func main() {
	// connecting to nats streaming
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

}
