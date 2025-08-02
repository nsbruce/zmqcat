package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pebbe/zmq4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zmqcat <protocol://ip:port>")
		os.Exit(1)
	}
	host := os.Args[1]

	context, err := zmq4.NewContext()
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	defer context.Term()

	subscriber, err := context.NewSocket(zmq4.SUB)
	if err != nil {
		log.Fatalf("Failed to create subscriber socket: %v", err)
	}
	defer subscriber.Close()

	err = subscriber.Connect(host)
	if err != nil {
		log.Fatalf("Failed to connect subscriber to %v: %v", host, err)
	}

	err = subscriber.SetSubscribe("")
	if err != nil {
		log.Fatalf("Failed to subscribe to all topics: %v", err)
	}

	fmt.Println("Waiting for messages")

	for {
		msg, meta, err := subscriber.RecvBytesWithMetadata(0)
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			continue
		}
		log.Println("Received message: length", len(msg), "bytes,", "metadata:", meta)
	}
}
