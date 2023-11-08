package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	topic := "contoh-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("Failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("Satu")},
		kafka.Message{Value: []byte("Dua")},
		kafka.Message{Value: []byte("Tiga")},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
