package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	reader *kafka.Reader
}

func NewKafkaReader(r *kafka.Reader) *KafkaReader {
	return &KafkaReader{reader: r}
}

func (r *KafkaReader) ReadMessage(ctx context.Context) ([]byte, error) {
	data, err := r.reader.ReadMessage(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data.Value, nil
}

func (r *KafkaReader) Close() error {
	err := r.reader.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
