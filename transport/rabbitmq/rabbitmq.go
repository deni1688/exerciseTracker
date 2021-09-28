package rabbitmq

import (
	"deni1688/exercise_tracker/domain"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

const CreatedExercise = "created_exercises"
type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal("error initializing broker connection", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("error initializing broker channel", err)
	}

	if err = ch.ExchangeDeclare(
		"exercise_events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatal("error initializing broker exchange", err)
	}

	return &RabbitMQ{conn, ch}
}

func (r *RabbitMQ) PublishCreated(ex *domain.Exercise) error {
	return r.publish(ex, CreatedExercise)
}

func (r *RabbitMQ) publish(ex *domain.Exercise, exerciseEvent string) error {
	bt, _ := json.Marshal(ex)

	return r.channel.Publish(
		"exercise_events",
		exerciseEvent,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bt,
		},
	)

}

func (r *RabbitMQ) Close() {
	_ = r.connection.Close()
	_ = r.channel.Close()
}
