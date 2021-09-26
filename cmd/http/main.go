package main

import (
	"deni1688/exercise_tracker/config"
	"deni1688/exercise_tracker/domain"
	"deni1688/exercise_tracker/storage"
	"deni1688/exercise_tracker/transport/http"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	env := config.GetFlags()

	err := config.LoadConfigFile(env)
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	er, err := storage.NewExerciseRepository(config.GetString("storage.driver"))
	if err != nil {
		log.Fatal("error creating new repository:", err)
	}

	es := domain.NewExerciseService(er)

	router := gin.Default()
	http.NewResource(router, http.GetHandlerFor(es, domain.ExerciseCollection))

	log.Fatal("error starting server:", router.Run(config.GetString("server.port")))
}
