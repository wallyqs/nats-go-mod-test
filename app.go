package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}

	doneCh := make(chan struct{})
	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("[Received]", string(msg.Data))
		doneCh <- struct{}{}
	})
	nc.Publish("foo", []byte("Hello Go modules!"))

	<-doneCh
}
