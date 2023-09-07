package supervisor

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func logf(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	fmt.Println()
}

type workerFunc func(topic string, partition int, offset int64, headers []kafka.Header, key, value []byte) (string, error)

type EventWorker struct {
	Name              string
	KafkaReaderConfig kafka.ReaderConfig
	WorkerFunc        workerFunc
}

func NewEventWorker(name string, kc kafka.ReaderConfig, wf workerFunc) *EventWorker {
	return &EventWorker{
		Name:              name,
		KafkaReaderConfig: kc,
		WorkerFunc:        wf,
	}
}

func (w *EventWorker) Serve(ctx context.Context) error {
	reader := kafka.NewReader(w.KafkaReaderConfig)

	for {
		m, err := reader.FetchMessage(ctx)

		if err != nil {
			panic(err.Error())
		}

		status, err := w.WorkerFunc(m.Topic, m.Partition, m.Offset, m.Headers, m.Key, m.Value)

		if err != nil {
			panic(err.Error())
		}

		if status == "SUCCESS" || status == "SKIP" {
			if err := reader.CommitMessages(ctx, m); err != nil {
				log.Fatal("failed to commit messages:", err)
			}
		}
	}
}
