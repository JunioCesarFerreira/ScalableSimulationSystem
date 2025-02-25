package kafkaclient

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer cria um novo consumidor Kafka
func NewKafkaConsumer(broker string, topic string) (*KafkaConsumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	return &KafkaConsumer{reader: reader}, nil
}

// ConsumeMessage consome uma mensagem do tópico Kafka
func (kc *KafkaConsumer) ConsumeMessage(ctx context.Context) (kafka.Message, error) {
	msg, err := kc.reader.ReadMessage(ctx)
	if err != nil {
		return kafka.Message{}, fmt.Errorf("failed to read message: %v", err)
	}
	return msg, nil
}

// Close fecha o consumidor Kafka
func (kc *KafkaConsumer) Close() error {
	if err := kc.reader.Close(); err != nil {
		return fmt.Errorf("failed to close Kafka reader: %v", err)
	}
	return nil
}
