package rest

import (
	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	*defaultController
}

func (ct *exerciseController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateExercise not implmented",
	})
}

func (ct *exerciseController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllExercises not implmented",
	})
}

type weightController struct {
	*defaultController
}

func (ct *weightController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateWeight not implmented",
	})
}

func (ct *weightController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllWeight not implmented",
	})
}

type caloriesController struct {
	*defaultController
}

func (ct *caloriesController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateCalories not implmented",
	})
}

func (ct *caloriesController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllCalories not implmented",
	})
}
