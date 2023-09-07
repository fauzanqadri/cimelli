package userinterestedonupdate

import (
	"consumers/supervisor"
	"encoding/json"
	"fmt"
	"producers"
	"time"

	"github.com/segmentio/kafka-go"
)

func worker(topic string, partition int, offset int64, headers []kafka.Header, key, value []byte) (string, error) {
	var userMessage producers.UserMessage

	if err := json.Unmarshal(value, &userMessage); err != nil {
		panic(err.Error())
	}

	if userMessage.Meta.Action == "UPDATE" {
		fmt.Printf("proccesing %v UPDATE ACTION with user id: %v, partition: %v, offset: %v\n", topic, userMessage.UserId, partition, offset)
		time.Sleep(500 * time.Millisecond) // simulate 500ms latency proccess
		fmt.Printf("processed %v UPDATE ACTION with user id: %v, partition: %v, offset: %v\n", topic, userMessage.UserId, partition, offset)
		return "SUCCESS", nil

	} else {
		fmt.Printf("skip %v non UPDATE ACTION with user id: %v, partition: %v, offset: %v\n", topic, userMessage.UserId, partition, offset)
		return "SKIP", nil
	}
}

func logf(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	fmt.Println()
}

func NewSuperVisor() *supervisor.EventWorker {
	config := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "consumer-user-interested-on-update-simple-supervisor-v1",
		Topic:       "user-event-v1",
		MinBytes:    5,
		MaxBytes:    10e6,
		Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
		StartOffset: kafka.LastOffset,
	}

	return supervisor.NewEventWorker("UserInterestedOnDelete", config, worker)
}
