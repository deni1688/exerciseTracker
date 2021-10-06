package mykafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/deni1688/exercise_tracker/domain"
	"log"
)

type KafkaProducer struct {
	p *kafka.Producer
}

func NewKafkaProducer(server string) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		log.Fatalln("error creating new kafka producer")
	}

	return &KafkaProducer{p}
}

func (b *KafkaProducer) PublishCreated(ex *domain.Exercise) error {
	return b.publish(ex, "exercise.created")
}

func (b *KafkaProducer) PublishUpdated(ex *domain.Exercise) error {
	return b.publish(ex, "exercise.updated")
}

func (b *KafkaProducer) publish(ex *domain.Exercise, topic string) error {
	body, _ := json.Marshal(ex)
	if err := b.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          body,
	}, nil); err != nil {
		return err
	}
	return nil
}

func (b *KafkaProducer) Close() {
	b.p.Flush(15 * 1000)
	b.p.Close()
}
