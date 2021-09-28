package main

import (
	"deni1688/exercise_tracker/config"
	"deni1688/exercise_tracker/domain"
	"deni1688/exercise_tracker/transport/rabbitmq"
	"encoding/json"
	"log"
)

func main() {
	env := config.GetEnv()
	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	consumer := rabbitmq.NewConsumer()
	msgs, err := consumer.GetMessages()
	if err != nil {
		log.Fatal("error consuming queue messages:", err)
	}

	forever := make(chan bool)
	var ex *domain.Exercise
	go func() {
		for d := range msgs {
			_ = json.Unmarshal(d.Body, &ex)
			log.Printf("New Exercise event of type %s with id %s", d.RoutingKey, ex.ID)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
