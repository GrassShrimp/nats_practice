package main

import (
	"flag"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
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

	defer func(_nc *nats.Conn) {
		_nc.Close()
	}(nc)

	for {
		// Simple Publisher
		err = nc.Publish("foo", []byte("Hello World"))

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(5) * time.Second)
	}

}
