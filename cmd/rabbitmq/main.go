package main

import (
	"encoding/json"
	"github.com/deni1688/exercise_tracker/config"
	"github.com/deni1688/exercise_tracker/domain"
	"github.com/deni1688/exercise_tracker/rabbitmq"
	"log"
)

func main() {
	env := config.GetEnv()

	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	conn, ch, err := rabbitmq.Connect()
	if err != nil {
		log.Fatal("error initializing connection or channel:", err)
	}

	messages, err := rabbitmq.NewConsumer(conn, ch).GetMessages(rabbitmq.CreatedExercise)
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
