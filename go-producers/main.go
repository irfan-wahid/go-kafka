package main

import (
	"time"

	appKafka "go-producers/kafka"
)

func main() {
	go appKafka.StartKafka()

	time.Sleep(10 * time.Minute)
}
