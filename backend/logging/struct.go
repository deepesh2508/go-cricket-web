package logging

import (
	"context"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

type kafkaLogger struct {
	kafka.Writer
}

func (k *kafkaLogger) Write(message []byte) (n int, err error) {
	err = k.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		})
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return len(message), nil
}
