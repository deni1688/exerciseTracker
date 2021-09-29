package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewConsumer(conn *amqp.Connection, ch *amqp.Channel) *Consumer {
	return &Consumer{conn, ch}
}

func (c *Consumer) GetMessages(exerciseEvent string) (<-chan amqp.Delivery, error) {
	q, err := c.channel.QueueDeclare(
		"",
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
		q.Name,
		exerciseEvent,
		exerciseExchange,
		false,
		nil)
	if err != nil {
		log.Fatal("error binding a queue", err)
	}

	return c.channel.Consume(
		q.Name,
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
