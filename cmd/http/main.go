package main

import (
	"deni1688/exercise_tracker/config"
	"deni1688/exercise_tracker/domain"
	http2 "deni1688/exercise_tracker/http"
	rabbitmq2 "deni1688/exercise_tracker/rabbitmq"
	"deni1688/exercise_tracker/storage"
	"github.com/gin-gonic/gin"
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

	conn, ch, err := rabbitmq2.Connect()
	if err != nil {
		log.Fatal("error initializing connection or channel:", err)
	}

	br := rabbitmq2.NewProducer(conn, ch)
	defer br.Close()

	es := domain.NewExerciseService(er, br)

	router := gin.Default()
	http2.NewResource(router, http2.GetHandlerFor(es, domain.ExerciseCollection))

	log.Fatal("error starting server:", router.Run(config.GetString("server.port")))
}
