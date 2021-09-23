package main

import (
	"deni1688/exerciseTracker/config"
	"deni1688/exerciseTracker/domain"
	"deni1688/exerciseTracker/rest"
	"deni1688/exerciseTracker/storage"
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
	controller := rest.NewControllerFactory(srv)

	router := gin.Default()
	rest.NewResource(router, controller.For(domain.Collection))

	log.Fatal("error starting server:", router.Run(config.GetString("server.port")))
}
