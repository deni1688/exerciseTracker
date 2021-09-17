package rest

import (
	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	*defaultController
}

func (h *exerciseController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateExercise not implmented",
	})
}

func (h *exerciseController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllExercises not implmented",
	})
}

type weightController struct {
	*defaultController
}

func (h *weightController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateWeight not implmented",
	})
}

func (h *weightController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllWeight not implmented",
	})
}

type caloriesController struct {
	*defaultController
}

func (h *caloriesController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateCalories not implmented",
	})
}

func (h *caloriesController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllCalories not implmented",
	})
}
