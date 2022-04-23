package sender

import (
	"log"

	"github.com/nats-io/nats.go"
)

func SendData(data string) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	nc.Publish("Shitposting", []byte(data))
}
