package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Printf("hello, world\n")

	writeToKafka()

	fmt.Printf("written\n")
}

func writeToKafka() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "quickstart-events",
		Balancer: &kafka.LeastBytes{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
		Key: []byte("Key-A"),
		Value: []byte("Hello from Go!"),
	})

	w.Close()
}
