package rest

import (
	"deni1688/exerciseTracker/domain"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	*defaultController
}

func (ct *exerciseController) Create(c *gin.Context) {
	var ex *domain.Exercise

	err := json.NewDecoder(c.Request.Body).Decode(&ex)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := ct.service.SaveExercise(ex)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func (ct *exerciseController) GetAll(c *gin.Context) {
	results, err := ct.service.ListExercises()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data":  results,
		"total": len(*results),
	})
}
