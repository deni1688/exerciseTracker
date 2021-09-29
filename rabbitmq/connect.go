package rabbitmq

import (
	"github.com/deni1688/exercise_tracker/config"
	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(config.GetString("broker.uri"))
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}
