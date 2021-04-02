package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("foo")
	nats.Connect("demo.nats.io")
}
