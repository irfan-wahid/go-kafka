package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "contoh-topic",
		MaxBytes: 10,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}
		fmt.Println("Message is:", string(m.Value))

	}
}
