package main

import (
	"deni1688/myHealthTrack/rest"
	"deni1688/myHealthTrack/storage"
	"deni1688/myHealthTrack/tracker"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	storageType := os.Getenv("MHT_STORAGE")
	repo, err := storage.RepositoryFactory(storageType)
	if err != nil {
		log.Fatal(err.Error())
	}
	trackerService := tracker.NewTrackerService(repo)
	controllerFactory := rest.NewControllerFactory(trackerService)

	rest.NewResource(router, tracker.EXERCISE).With(controllerFactory.Create(tracker.EXERCISE))
	rest.NewResource(router, tracker.WEIGHT).With(controllerFactory.Create(tracker.WEIGHT))

	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal("Error running server", err)
	}
}
