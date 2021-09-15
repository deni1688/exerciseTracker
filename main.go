package main

import "github.com/gin-gonic/gin"

func main() {
	newApi(gin.Default()).
		InitExerciseApi(GetHandlersFor("exercises")).
		InitWeightApi(GetHandlersFor("weight")).
		InitNutritionApi(GetHandlersFor("nutrition")).
		Run()
}
