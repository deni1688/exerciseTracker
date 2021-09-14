package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	newApi(r).
		InitExerciseApi().
		InitWeightApi().
		Run()
}
