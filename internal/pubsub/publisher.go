package pubsub

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	sc stan.Conn
}

func (p *Publisher) Start() error {
	//defer p.sc.Close()
	for {
		data, err := json.Marshal(randomOrder())
		if err != nil {
			return fmt.Errorf("publisher: cannot marshal msg to json")
		}
		p.sc.Publish("test1", data)
		time.Sleep(1 * time.Second)
		log.Print("Publisher send the message")
	}
}

func (p *Publisher) SendMessage() error {

	return nil
}

func NewPublisher(sc stan.Conn) *Publisher {
	return &Publisher{
		sc: sc,
	}
}
