package rest

import (
	"deni1688/myHealthTrack/domain"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	HandleCreate(c *gin.Context)
	HandleGetAll(c *gin.Context)
	HandleGetOne(c *gin.Context)
}

type defaultHandlers struct {
	service domain.Service
}

func (h *defaultHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreate not implmented",
	})
}

func (h *defaultHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAll not implmented",
	})
}

func (h *defaultHandlers) HandleGetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetOne not implmented",
	})
}

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

func GetHandlersFor(handlers string, service domain.Service) Handlers {
	switch handlers {
	case "exercises":
		return &exerciseHandlers{&defaultHandlers{service}}
	case "weight":
		return &weightHandlers{&defaultHandlers{service}}
	case "nutrition":
		return &nutritionHandlers{&defaultHandlers{service}}
	default:
		return &defaultHandlers{service}
	}
}
