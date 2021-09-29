package http

import (
	"deni1688/exercise_tracker/domain"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type exerciseHandler struct {
	service  domain.ExerciseService
	resource string
}

func (h *exerciseHandler) Create(c *gin.Context) {
	var ex *domain.ExerciseRequest

	err := json.NewDecoder(c.Request.Body).Decode(&ex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.SaveExercise(ex)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func (h *exerciseHandler) GetAll(c *gin.Context) {
	results, err := h.service.ListExercises()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  results,
		"total": len(*results),
	})
}

func (h *exerciseHandler) GetResource() string {
	return h.resource
}
