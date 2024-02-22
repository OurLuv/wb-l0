package pubsub

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	sc stan.Conn
}

func (p *Publisher) Start() error {
	for {
		p.SendMessage()
	}
}

func (p *Publisher) SendMessage() error {
	data, err := json.Marshal(RandomOrder())
	if err != nil {
		return err
	}
	err = p.sc.Publish("test1", data)
	if err != nil {
		log.Printf("Error publishing message: %v", err)
	}
	time.Sleep(20 * time.Second)
	log.Print("Publisher send the message")
	return nil
}

func NewPublisher(sc stan.Conn) *Publisher {
	return &Publisher{
		sc: sc,
	}
}
