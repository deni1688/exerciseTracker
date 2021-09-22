package main

import (
	"deni1688/exerciseTracker/config"
	"deni1688/exerciseTracker/domain"
	"deni1688/exerciseTracker/http"
	"deni1688/exerciseTracker/storage"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := config.LoadConfig(getEnv())
	if err != nil {
		log.Fatal("error loading the config file:", err)
	}

	repo, err := storage.NewExerciseRepository("sqlite")
	if err != nil {
		log.Fatal("error creating new repository:", err)
	}

	srv := domain.NewExerciseService(repo)
	controller := http.NewControllerFactory(srv)

	router := gin.Default()
	http.NewResource(router, controller.For(domain.Collection))

	log.Fatal("error starting server:", router.Run(config.GetString("api.port")))
}

func getEnv() *string {
	env := flag.String("environment", "testing", "sets the config environments")
	flag.Parse()
	return env
}
