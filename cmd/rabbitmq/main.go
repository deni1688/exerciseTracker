package main

import (
	"deni1688/exercise_tracker/config"
	"deni1688/exercise_tracker/domain"
	rabbitmq2 "deni1688/exercise_tracker/rabbitmq"
	"encoding/json"
	"log"
)

func main() {
	env := config.GetEnv()

	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	conn, ch, err := rabbitmq2.Connect()
	if err != nil {
		log.Fatal("error initializing connection or channel:", err)
	}

	messages, err := rabbitmq2.NewConsumer(conn, ch).
		GetMessages(rabbitmq2.CreatedExercise)
	if err != nil {
		log.Fatal("error consuming queue messages:", err)
	}

	ex := new(domain.Exercise)
	forever := make(chan bool)

	go func() {
		for d := range messages {
			_ = json.Unmarshal(d.Body, &ex)
			log.Printf("New Exercise event of type %s with id %s", d.RoutingKey, ex.ID)
		}
	}()

	log.Printf(" [*] Waiting for exercise events. To exit press CTRL+C")
	<-forever
}
