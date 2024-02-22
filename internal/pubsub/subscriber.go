package pubsub

import (
	"encoding/json"
	"log"

	"github.com/OurLuv/l0/internal/model"
	"github.com/OurLuv/l0/internal/service"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	sc      stan.Conn
	service service.OrderServcie
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
		log.Print(err)
	}
	_, err = s.service.Save(order)
	if err != nil {
		log.Print(err)
	}
	return nil
}

func NewSubscriber(sc stan.Conn, service service.OrderServcie) *Subscriber {
	return &Subscriber{
		sc:      sc,
		service: service,
	}
}
