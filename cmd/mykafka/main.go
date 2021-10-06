package main

import (
	"encoding/json"
	"github.com/deni1688/exercise_tracker/config"
	"github.com/deni1688/exercise_tracker/domain"
	"github.com/deni1688/exercise_tracker/mykafka"
	"log"
)

func main() {
	env := config.GetEnv()

	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	c := mykafka.NewKafkaConsumer(config.GetString("brokers.kafka.server"))

	messages := c.GetMessages()

	ex := new(domain.Exercise)
	forever := make(chan bool)

	go func() {
		for d := range messages {
			_ = json.Unmarshal(d.Value, &ex)
			log.Printf("New Exercise event of type %s with id %s", d.String(), ex.ID)
		}
	}()

	log.Printf(" [*] Waiting for exercise events. To exit press CTRL+C")
	<-forever
}
