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

	ts := tracker.NewTrackerService()
	cf := rest.NewControllerFactory(ts)
	r := gin.Default()
	rf := rest.NewRouteFactory(r)

	rf.Create(tracker.EXERCISE, cf.Create(tracker.EXERCISE))
	rf.Create(tracker.WEIGHT, cf.Create(tracker.WEIGHT))
	rf.Create(tracker.CALORIES, cf.Create(tracker.CALORIES))

	r.Run(os.Getenv("PORT"))
}
