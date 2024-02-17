package pubsub

import (
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	sc stan.Conn
}

func (p *Publisher) Start() error {
	defer p.sc.Close()
	for {
		p.sc.Publish("test1", []byte("first tests"))
		time.Sleep(1 * time.Second)
		log.Print("Publisher send the message")
	}
}

func NewPublisher(sc stan.Conn) *Publisher {
	return &Publisher{
		sc: sc,
	}
}
