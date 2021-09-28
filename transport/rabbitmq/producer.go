package rabbitmq

import (
	"deni1688/exercise_tracker/config"
	"deni1688/exercise_tracker/domain"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

const (
	UpdatedExercise  = "created_exercises"
	CreatedExercise  = "updated_exercises"
	ExerciseExchange = "exercise_events"
)

type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewProducer() *Producer {
	conn, err := amqp.Dial(config.GetString("broker.uri"))
	if err != nil {
		log.Fatal("error initializing broker connection", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("error initializing broker channel", err)
	}

	if err = ch.ExchangeDeclare(
		ExerciseExchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatal("error initializing broker exchange", err)
	}

	return &Producer{conn, ch}
}

func (p *Producer) PublishCreated(ex *domain.Exercise) error {
	return p.publish(ex, CreatedExercise)
}
func (p *Producer) PublishUpdated(ex *domain.Exercise) error {
	return p.publish(ex, UpdatedExercise)
}
func (p *Producer) publish(ex *domain.Exercise, exerciseEvent string) error {
	bt, _ := json.Marshal(ex)

	return p.channel.Publish(
		ExerciseExchange,
		exerciseEvent,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bt,
		},
	)

}

func (p *Producer) Close() {
	_ = p.connection.Close()
	_ = p.channel.Close()
}
