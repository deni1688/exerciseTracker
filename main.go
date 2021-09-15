package main

import (
	"deni1688/myHealthTrack/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	rest.New(gin.Default()).
		InitExerciseRoutes(rest.GetHandlersFor("exercises")).
		InitWeightRoutes(rest.GetHandlersFor("weight")).
		InitNutritionRoutes(rest.GetHandlersFor("nutrition")).
		Run()
}
