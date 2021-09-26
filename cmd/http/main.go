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

	repo, err := storage.NewExerciseRepository(config.GetString("storage.driver"))
	if err != nil {
		log.Fatal("error creating new repository:", err)
	}

	srv := domain.NewExerciseService(repo)
	handler := http.NewHandlerFactory(srv)

	router := gin.Default()
	http.NewResource(router, handler.For(domain.Collection))

	log.Fatal("error starting server:", router.Run(config.GetString("server.port")))
}
