package main

import (
	"deni1688/myHealthTrack/tracker"
	"deni1688/myHealthTrack/rest"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	s := tracker.NewService()

	if err := rest.New(gin.Default()).
		InitExerciseRoutes(rest.CreateControllerFor("exercise", s)).
		InitWeightRoutes(rest.CreateControllerFor("weight", s)).
		InitCaloriesRoutes(rest.CreateControllerFor("calories", s)).
		Run(":9090"); err != nil {
		log.Fatal("Error starting the rest service:", err)
	}

}
