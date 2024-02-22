package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OurLuv/l0/internal/handler"
	"github.com/OurLuv/l0/internal/pubsub"
	"github.com/OurLuv/l0/internal/service"
	"github.com/OurLuv/l0/internal/storage/cache"
	"github.com/OurLuv/l0/internal/storage/postgres"
	"github.com/nats-io/stan.go"
)

func main() {
	//* storage
	pool, err := postgres.NewPool("postgres://postgres:admin@localhost:5432/wbl0")
	if err != nil {
		log.Printf("failed to init storage: %s", err)
		os.Exit(1)
	}
	defer pool.Close()
	log.Printf("Storage init")
	//creating repo
	repo := postgres.NewOrderRepository(pool)

	//* service
	service := service.New(repo, cache.New())
	//pulling data from DB to cache
	if err := service.Pull(); err != nil {
		fmt.Printf("can't pull data from DB: %s", err.Error())
	}

	//* handler
	h := handler.NewHandler(service)

	//init routes
	r := h.InitRoutes()

	//creating server
	server := handler.NewServer(r)

	//starting server
	go func() {
		log.Printf("Server started")
		if err := server.Start(); err != nil {
			log.Fatalf("can't start a server: %s", err.Error())
		}
	}()

	//* nats streaming
	sc, err := stan.Connect("my_cluster", "cl1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// creating and starting new subscriber
	s := pubsub.NewSubscriber(sc, service)
	s.Start()
	// creating and starting new subscriber
	p := pubsub.NewPublisher(sc)
	p.Start()

}
