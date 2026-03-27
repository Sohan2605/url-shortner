package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func Init() {
	Writer = &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "click-events",
	}
}

func Publish(code string) {
	Writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(code)},
	)
}
