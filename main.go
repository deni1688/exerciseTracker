package main

import (
	"deni1688/myHealthTrack/domain"
	"deni1688/myHealthTrack/rest"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	s := domain.NewService()

	if err := rest.New(gin.Default()).
		InitExerciseRoutes(rest.GetHandlersFor("exercise", s)).
		InitWeightRoutes(rest.GetHandlersFor("weight", s)).
		InitNutritionRoutes(rest.GetHandlersFor("nutrition", s)).
		Run(":9090"); err != nil {
		log.Fatal("Error starting the rest service:", err)
	}

}
