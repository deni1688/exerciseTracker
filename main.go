package main

import (
	"deni1688/myHealthTrack/rest"
	"deni1688/myHealthTrack/storage"
	"deni1688/myHealthTrack/tracker"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file", err)
	}

	repo, err := storage.NewRepository()
	if err != nil {
		log.Fatal("Error creating new repository", err)
	}

	trackerService := tracker.NewTrackerService(repo)
	controllerFactory := rest.NewControllerFactory(trackerService)

	router := gin.Default()
	rest.NewResource(router, tracker.EXERCISE).With(controllerFactory.Create(tracker.EXERCISE))
	rest.NewResource(router, tracker.WEIGHT).With(controllerFactory.Create(tracker.WEIGHT))

	port := os.Getenv("PORT")
	log.Fatal(router.Run(port))
}
