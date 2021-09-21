package main

import (
	"deni1688/exerciseTracker/exercises"
	"deni1688/exerciseTracker/http/rest"
	"deni1688/exerciseTracker/http/rest/controllers"
	"deni1688/exerciseTracker/repository"
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
	controller := controllers.NewControllerFactory(trackerService)

	router := gin.Default()
	rest.NewResource(router, controller.For(exercises.Collection))

	port := os.Getenv("PORT")
	log.Fatal(router.Run(port))
}
