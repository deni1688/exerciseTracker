package main

import (
	"deni1688/myHealthTrack/rest"
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

	trackingService := tracker.NewTrackerService()
	controllerFactory := rest.NewControllerFactory(trackingService)
	resourceFactory := rest.NewResourceFactory(router)

	resourceFactory.Create(tracker.EXERCISE, controllerFactory.Create(tracker.EXERCISE))
	resourceFactory.Create(tracker.WEIGHT, controllerFactory.Create(tracker.WEIGHT))
	resourceFactory.Create(tracker.CALORIES, controllerFactory.Create(tracker.CALORIES))



	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal("Error running server", err)
	}

}
