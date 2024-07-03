package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	// In the `jetstream` package, almost all API calls rely on `context.Context` for timeout/cancellation handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nc, _ := nats.Connect(nats.DefaultURL)

	// Create a JetStream management interface
	js, _ := jetstream.New(nc)

	// Create a stream
	s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "YOUTUBE",
		Subjects: []string{"YOUTUBE.*"},
	})

	// Publish some messages
	js.Publish(ctx, "YOUTUBE.url", []byte("Wcmpl2596G0"))

	// Create durable consumer
	c, _ := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "YOUTUBE",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	// Get 10 messages from the consumer
	messageCounter := 0
	msgs, err := c.Fetch(10)
	if err != nil {
		// handle error
	}

	for msg := range msgs.Messages() {
		msg.Ack()
		fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))
		messageCounter++
	}
}
