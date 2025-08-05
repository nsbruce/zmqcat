package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pebbe/zmq4"
	"github.com/urfave/cli/v3"
)

var version = "dev" // will be set by GoReleaser on build

func main() {
	var verbose bool
	var host string
	cmd := &cli.Command{
		Name:      "zmqcat",
		Usage:     "inspect published ZeroMQ messages",
		ArgsUsage: "<protocol>://<host>:<port>",
		Version:   version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Value:       false,
				Usage:       "Print message and metadata contents",
				Destination: &verbose,
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "host",
				UsageText:   "The ZeroMQ host to connect to, in the format <protocol>://<host>:<port>",
				Destination: &host,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if host == "" {
				cli.ShowAppHelp(cmd)
				return fmt.Errorf("missing required argument")
			}
			process_connection(host, verbose)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func process_connection(host string, verbose bool) {

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
		// msg, meta, err := subscriber.RecvMessageWithMetadata(0) // 0 means default options
		msg, meta, err := subscriber.RecvBytesWithMetadata(0) // 0 means default options
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			continue
		}
		if verbose {
			if len(msg) > 1024 {
				log.Println("Received message of length", len(msg), "bytes (too long to print).")
			} else {
				log.Println("Received message:", string(msg))
			}
		} else {
			log.Println("Received message of length", len(msg), "bytes.")
		}
		if verbose && len(meta) > 0 {
			fmt.Println("Metadata:")
			for key, value := range meta {
				fmt.Println("   ", key, "=", value)
			}
		}
	}
}
