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
		InitExerciseRoutes(rest.CreateHandlersFor("exercise", s)).
		InitWeightRoutes(rest.CreateHandlersFor("weight", s)).
		InitNutritionRoutes(rest.CreateHandlersFor("nutrition", s)).
		Run(":9090"); err != nil {
		log.Fatal("Error starting the rest service:", err)
	}

}
