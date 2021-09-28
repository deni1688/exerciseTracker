package rabbitmq

import (
	"deni1688/exercise_tracker/config"
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewConsumer() *Consumer {
	conn, err := amqp.Dial(config.GetString("broker.uri"))
	if err != nil {
		log.Fatal("error connecting to RabbitMQProducer", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("error opening a channel", err)
	}

	err = ch.ExchangeDeclare(
		ExerciseExchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("error initializing an exchange", err)
	}

	_, err = ch.QueueDeclare(
		CreatedExercise,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("error initializing a queue")
	}

	err = ch.QueueBind(
		CreatedExercise,
		CreatedExercise,
		ExerciseExchange,
		false,
		nil)
	if err != nil {
		log.Fatal("error binding a queue", err)
	}

	return &Consumer{conn, ch}
}
func (c *Consumer) GetMessages() (<-chan amqp.Delivery, error) {
	return c.channel.Consume(
		CreatedExercise,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (c *Consumer) Close() {
	_ = c.connection.Close()
	_ = c.channel.Close()
}
