package main

import "github.com/gin-gonic/gin"

type Handlers interface {
	HandleCreate(c *gin.Context)
	HandleGetAll(c *gin.Context)
	HandleGetOne(c *gin.Context)
}

type DefaultHandlers struct {
}

func (h *DefaultHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreate not implmented",
	})
}

func (h *DefaultHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAll not implmented",
	})
}

func (h *DefaultHandlers) HandleGetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetOne not implmented",
	})
}

type exerciseHandlers struct {
	*DefaultHandlers
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
	*DefaultHandlers
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
	*DefaultHandlers
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
