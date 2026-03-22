package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "blockchain_events",
		GroupID: "event_workers",
	})

	log.Println("Worker3 started...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Worker3 error fetching message:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Worker3 received:", string(msg.Value))
	}
}
