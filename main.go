package main

import (
	"deni1688/myHealthTrack/http/rest"
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
	rest.NewResource(router, tracker.Collection).With(controllerFactory.Create(tracker.Collection))

	port := os.Getenv("PORT")
	log.Fatal(router.Run(port))
}
