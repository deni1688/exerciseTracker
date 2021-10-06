package main

import (
	"github.com/deni1688/exercise_tracker/cli"
	"github.com/deni1688/exercise_tracker/config"
	"github.com/deni1688/exercise_tracker/domain"
	"github.com/deni1688/exercise_tracker/rabbitmq"
	"github.com/deni1688/exercise_tracker/storage"
	"log"
)

func main() {
	env := config.GetEnv()

	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	er, err := storage.NewExerciseRepository(config.GetString("storage.driver"))
	if err != nil {
		log.Fatal("error creating new repository:", err)
	}

	conn, ch, err := rabbitmq.Connect()
	if err != nil {
		log.Fatal("error initializing connection or channel:", err)
	}

	br := rabbitmq.NewRabbitMQProducer(conn, ch)
	defer br.Close()
	es := domain.NewExerciseService(er, br)

	if err := cli.NewExerciseCLI(es).Execute(); err != nil {
		log.Fatal("error executing cli")
	}
}
