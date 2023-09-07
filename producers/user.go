package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"models"

	kafka "github.com/segmentio/kafka-go"
)

type Meta struct {
	EventId        string `json:"event_id"`
	EventTimestamp string `json:"event_timestamp"`
	Action         string `json:"action"`
}

type UserMessage struct {
	UserId string       `json:"user_id"`
	Meta   Meta         `json:"meta"`
	Before *models.User `json:"before"`
	After  *models.User `json:"after"`
}

var w *kafka.Writer

func init() {
	w = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Balancer: &kafka.Hash{},
	}
}

func PublishUserCreation(after *models.User) {
	eventId, err := models.Sf.NextID()
	userId := fmt.Sprintf("%v", after.Id)

	if err != nil {
		fmt.Println(err)
		return
	}

	meta := &Meta{
		EventId:        fmt.Sprintf("%v", eventId),
		EventTimestamp: time.Now().Format(time.RFC3339),
		Action:         "CREATE",
	}

	message := &UserMessage{
		UserId: userId,
		After:  after,
		Meta:   *meta,
	}

	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "user-event-v1",
			Key:   []byte(userId),
			Value: msg,
		},
	)
}

func PublishUserDeletion(before *models.User) {
	eventId, err := models.Sf.NextID()
	userId := fmt.Sprintf("%v", before.Id)

	if err != nil {
		fmt.Println(err)
		return
	}

	meta := &Meta{
		EventId:        fmt.Sprintf("%v", eventId),
		EventTimestamp: time.Now().Format(time.RFC3339),
		Action:         "DELETE",
	}

	message := &UserMessage{
		UserId: userId,
		Before: before,
		Meta:   *meta,
	}

	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "user-event-v1",
			Key:   []byte(userId),
			Value: msg,
		},
	)
}

func PublishUserUpdate(before, after *models.User) {
	eventId, err := models.Sf.NextID()
	userId := fmt.Sprintf("%v", after.Id)

	if err != nil {
		fmt.Println(err)
		return
	}

	meta := &Meta{
		EventId:        fmt.Sprintf("%v", eventId),
		EventTimestamp: time.Now().Format(time.RFC3339),
		Action:         "UPDATE",
	}

	message := &UserMessage{
		UserId: userId,
		After:  after,
		Before: before,
		Meta:   *meta,
	}

	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "user-event-v1",
			Key:   []byte(userId),
			Value: msg,
		},
	)
}
