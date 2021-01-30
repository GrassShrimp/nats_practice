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
		_nc.Drain()
		_nc.Close()
	}(nc)

	for {
		// Requests
		msg, err := nc.Request("help", []byte("help me"), time.Duration(10)*time.Second)

		if err != nil {
			continue
		}

		fmt.Println(string(msg.Data))

		time.Sleep(time.Duration(5) * time.Second)
	}

}
