package mykafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

const (
	UpdatedExercise string = "exercise.updated"
	CreatedExercise string = "exercise.created"
	exerciseGroupID string = "exercise.events"
)

type KafkaConsumer struct {
	c *kafka.Consumer
}

func NewKafkaConsumer(server string) *KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          exerciseGroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalln("error creating kafka consumer")
	}

	err = c.SubscribeTopics([]string{CreatedExercise, UpdatedExercise}, nil)
	if err != nil {
		log.Fatalln("error subscribing to topic")
	}

	return &KafkaConsumer{c}
}

func (b *KafkaConsumer) GetMessages() <-chan kafka.Message {
	msgs := make(chan kafka.Message)
	go func() {
		for {
			msg, err := b.c.ReadMessage(100 * time.Microsecond)
			if err == nil {
				msgs <- *msg
			}
		}
	}()
	return msgs
}

func (b *KafkaConsumer) Close() {
	_ = b.c.Close()
}
