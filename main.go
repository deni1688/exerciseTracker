package main

import (
	"deni1688/myHealthTrack/domain"
	"deni1688/myHealthTrack/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	s := domain.NewService()

	rest.New(gin.Default()).
		InitExerciseRoutes(rest.GetHandlersFor("exercises", s)).
		InitWeightRoutes(rest.GetHandlersFor("weight", s)).
		InitNutritionRoutes(rest.GetHandlersFor("nutrition", s)).
		Run(":9090")
}
