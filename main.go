package main

import (
	"deni1688/myHealthTrack/http/rest"
	"deni1688/myHealthTrack/repository"
	"deni1688/myHealthTrack/exercises"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file", err)
	}

	repo, err := repository.New()
	if err != nil {
		log.Fatal("Error creating new repository", err)
	}

	trackerService := exercises.NewService(repo)
	controllerFactory := rest.NewControllerFactory(trackerService)

	router := gin.Default()
	rest.NewResource(router, exercises.Collection).With(controllerFactory.Create(exercises.Collection))

	port := os.Getenv("PORT")
	log.Fatal(router.Run(port))
}
