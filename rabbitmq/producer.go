package rabbitmq

import (
	"deni1688/exercise_tracker/domain"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type ExerciseEventType string

const (
	UpdatedExercise  ExerciseEventType = "exercise.updated"
	CreatedExercise  ExerciseEventType = "exercise.created"
	exerciseExchange string = "exercise.events"
)

type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewProducer(conn *amqp.Connection, ch *amqp.Channel) *Producer {
	if err := ch.ExchangeDeclare(
		exerciseExchange,
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
func (p *Producer) publish(ex *domain.Exercise, ee ExerciseEventType) error {
	bt, _ := json.Marshal(ex)
	log.Println(ee)
	return p.channel.Publish(
		exerciseExchange,
		string(ee),
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
