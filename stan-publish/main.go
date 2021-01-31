package main

import (
	"flag"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

func main() {
	natsURL := flag.String("nats-url", nats.DefaultURL, "url of nats server")
	flag.Parse()
	// Connect to a server
	nc, err := nats.Connect(*natsURL)

	if err != nil {
		fmt.Printf("Try to connecting %s failure\n", *natsURL)
		panic(err)
	}

	sc, err := stan.Connect("demo", "demo-publish", stan.NatsConn(nc), stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
		fmt.Printf("Connection lost, reason: %v\n", reason)
	}))

	if err != nil {
		fmt.Println("Try to connecting nats streaming failure")
		panic(err)
	}

	defer func(_nc *nats.Conn, _sc stan.Conn) {
		_sc.Close()
		_nc.Close()
	}(nc, sc)

	for {
		// Simple Synchronous Publisher
		err := sc.Publish("foo", []byte("Hello World"))

		if err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}

}
