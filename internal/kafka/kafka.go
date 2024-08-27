package kafka

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaProducer *KafkaProducer
	kafkaConsumer *KafkaConsumer
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	Topic    string
}

func InitKafkaProducer() error {
	var (
		kafkaHost = os.Getenv("KAFKA_HOST")
	)

	config := kafka.ConfigMap{
		"bootstrap.servers": kafkaHost,
	}

	producer, err := kafka.NewProducer(&config)
	if err != nil {
		return err
	}

	kafkaProducer = &KafkaProducer{
		Producer: producer,
		Topic:    "account-message-events",
	}

	return nil
}

func InitKafkaConsumer() error {
	var (
		kafkaHost = os.Getenv("KAFKA_HOST")
	)

	config := kafka.ConfigMap{
		"bootstrap.servers": kafkaHost,
	}

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		return err
	}

	kafkaConsumer = &KafkaConsumer{
		Consumer: consumer,
		Topic:    "account-message-events",
	}

	return nil
}

func GetKafkaProducer() *KafkaProducer {
	if kafkaProducer == nil {
		InitKafkaProducer()
	}

	return kafkaProducer
}

func GetKafkaConsumer() *KafkaConsumer {
	if kafkaConsumer == nil {
		InitKafkaConsumer()
	}

	return kafkaConsumer
}
