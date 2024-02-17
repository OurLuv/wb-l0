package pubsub

import (
	"log"

	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	sc stan.Conn
}

func (s *Subscriber) Start() error {
	defer s.sc.Close()
	s.sc.Subscribe("test1", func(msg *stan.Msg) {
		log.Print("Subscriber recieved a message")
	})
	return nil
}

func (s *Subscriber) HandleMessage() {

}

func NewSubscriber(sc stan.Conn) *Publisher {
	return &Publisher{
		sc: sc,
	}
}
