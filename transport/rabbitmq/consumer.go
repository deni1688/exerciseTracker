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

	return &Consumer{conn, ch}
}

func (c *Consumer) GetMessages(exerciseEvent string) (<-chan amqp.Delivery, error) {
	_, err := c.channel.QueueDeclare(
		exerciseEvent,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("error initializing a queue")
	}

	err = c.channel.QueueBind(
		exerciseEvent,
		exerciseEvent,
		ExerciseExchange,
		false,
		nil)
	if err != nil {
		log.Fatal("error binding a queue", err)
	}
	return c.channel.Consume(
		exerciseEvent,
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
