package services

import (
	"bytes"
	"context"
	"encoding/json"

	db "tracking_system/internal/db/entities"
	r "tracking_system/internal/db/repositories"
	"tracking_system/internal/entities"
	k "tracking_system/internal/kafka"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type kafkaService struct {
	ctx             context.Context
	log             *logrus.Logger
	kafkaProducer   *k.KafkaProducer
	deliveryChannel chan kafka.Event

	accountEventRepo *r.AccountEventRepo
}

func newKafkaService(
	ctx context.Context,
	log *logrus.Logger,
	kafkaProducer *k.KafkaProducer,
	accountEventRepo *r.AccountEventRepo,
) *kafkaService {
	return &kafkaService{
		ctx,
		log,
		kafkaProducer,
		make(chan kafka.Event, 10000),
		accountEventRepo,
	}
}

func GetKafkaService(ctx context.Context, log *logrus.Logger) *kafkaService {
	return newKafkaService(
		ctx,
		log,
		k.GetKafkaProducer(),
		r.GetAccountEventRepo(ctx, log),
	)
}

// SendEvent sends event to kafka and logs it into database.
// Returns nil on success
func (s *kafkaService) SendEvent(event *entities.Event) error {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(event); err != nil {
		s.log.WithError(err).Error("Unable to encode data")
		return err
	}

	err := s.accountEventRepo.InsertAccountEvent((*db.AccountEvent)(event))
	if err != nil {
		s.log.WithError(err).Error("Unable to insert account event to db")
		return err
	}

	err = s.kafkaProducer.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &s.kafkaProducer.Topic,
			Partition: kafka.PartitionAny,
		},
		Value: b.Bytes()}, s.deliveryChannel)
	if err != nil {
		s.log.WithError(err).Error("Unable to send event to kafka")
		return err
	}

	<-s.deliveryChannel
	s.log.Info("Message sent!")

	return nil
}
