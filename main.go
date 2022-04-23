package main

import (
	"httpServices/sender"
	"httpServices/writer"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const data string = "1"

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	sub, err := nc.SubscribeSync("Shitposting")
	if err != nil {
		log.Fatal(err)
	}
	sender.SendData(data)
	msg, err := sub.NextMsg(99999 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("I've got something! ", msg.Data)
	file := writer.CreateFile("data.txt")
	writer.WriteTo(file, string(msg.Data))
}
