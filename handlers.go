package main

import "github.com/gin-gonic/gin"

type Handlers interface {
	HandleCreate(c *gin.Context)
	HandleGetAll(c *gin.Context)
}

func GetHandlersFor(handlers string) Handlers {
	switch handlers {
	case "exercises":
		return &exerciseHandlers{}
	case "weight":
		return &weightHandlers{}
	case "nutrition":
		return &nutritionHandlers{}
	default:
		return nil
	}
}

type exerciseHandlers struct {
}

func (eh *exerciseHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateExercise not implmented",
	})
}

func (eh *exerciseHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllExercises not implmented",
	})
}

type weightHandlers struct {
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
