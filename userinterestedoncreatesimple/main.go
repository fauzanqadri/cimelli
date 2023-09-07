package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"producers"

	kafka "github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func logf(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	fmt.Println()
}

func init() {
	config := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "consumer-user-interested-on-create-simple-v1",
		Topic:       "user-event-v1",
		MinBytes:    5,
		MaxBytes:    10e6,
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
		StartOffset: kafka.LastOffset,
	}

	reader = kafka.NewReader(config)
}

func main() {
	ctx := context.Background()

	for {
		m, err := reader.FetchMessage(ctx)

		if err != nil {
			break
		}

		var userMessage producers.UserMessage

		if err = json.Unmarshal(m.Value, &userMessage); err != nil {
			panic(err.Error())
		}

		if userMessage.Meta.Action == "CREATE" {
			fmt.Printf("proccesing %v with user id: %v, partition: %v, offset: %v\n", m.Topic, userMessage.UserId, m.Partition, m.Offset)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("processed %v with user id: %v, partition: %v, offset: %v\n", m.Topic, userMessage.UserId, m.Partition, m.Offset)
		} else {
			fmt.Printf("skip %v with user id: %v, partition: %v, offset: %v\n", m.Topic, userMessage.UserId, m.Partition, m.Offset)
		}

		if err := reader.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}
