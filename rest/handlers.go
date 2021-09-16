package rest

import (
	"github.com/gin-gonic/gin"
)

type exerciseHandlers struct {
	*defaultHandlers
}

func (h *exerciseHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateExercise not implmented",
	})
}

func (h *exerciseHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllExercises not implmented",
	})
}

type weightHandlers struct {
	*defaultHandlers
}

func (h *weightHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateWeight not implmented",
	})
}

func (h *weightHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllWeight not implmented",
	})
}

type nutritionHandlers struct {
	*defaultHandlers
}

func (h *nutritionHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateNutrition not implmented",
	})
}

func (h *nutritionHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllNutrition not implmented",
	})
}
