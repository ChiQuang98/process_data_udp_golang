package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

const (
	queue   = "Order.OrdersCreatedQueue"
	subject = "Order.OrderCreated"
)

func main() {
	//nc, _ := nats.Connect("10.16.150.132:4222")

	// Create server connection
	natsConnection, _ := nats.Connect("10.16.150.132:4222")
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message in Worker 1: %+v\n", eventStore)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
