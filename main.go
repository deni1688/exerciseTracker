package main

import (
	"deni1688/exerciseTracker/exercises"
	"deni1688/exerciseTracker/http/rest"
	"deni1688/exerciseTracker/sqlite"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file:", err)
	}

	repo, err := sqlite.NewExerciseRepository()
	if err != nil {
		log.Fatal("Error creating new repository:", err)
	}

	trackerService := exercises.NewExerciseService(repo)
	controller := rest.NewControllerFactory(trackerService)

	router := gin.Default()
	rest.NewResource(router, controller.For(exercises.Collection))

	port := os.Getenv("PORT")
	log.Fatal("Error starting server:", router.Run(port))
}
