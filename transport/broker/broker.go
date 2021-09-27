package broker

import (
	"deni1688/exercise_tracker/domain"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type rabbitMQBroker struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	qu    *amqp.Queue
	topic string
}

func NewBroker(topic string) domain.ExerciseBroker {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal("error initializing broker connection", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("error initializing broker channel", err)
	}

	qu, err := ch.QueueDeclare(
		domain.ExerciseCollection,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("error initializing broker queue", err)
	}

	return &rabbitMQBroker{conn, ch, &qu, topic}
}

func (b rabbitMQBroker) Publish(ex *domain.Exercise) error {
	bt, _ := json.Marshal(ex)

	return b.ch.Publish(
		"",
		b.topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bt,
		},
	)
}
