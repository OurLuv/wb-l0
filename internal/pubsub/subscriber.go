package pubsub

import (
	"encoding/json"
	"log"

	"github.com/OurLuv/l0/internal/model"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	sc stan.Conn
}

func (s *Subscriber) Start() error {
	//defer s.sc.Close()
	s.sc.Subscribe("test1", func(msg *stan.Msg) {
		s.HandleMessage(msg.Data)
	})
	return nil
}

func (s *Subscriber) HandleMessage(msg []byte) error {
	var order model.Order

	err := json.Unmarshal(msg, &order)
	if err != nil {
		return err
	}
	log.Print(order)
	return nil
}

func NewSubscriber(sc stan.Conn) *Subscriber {
	return &Subscriber{
		sc: sc,
	}
}
