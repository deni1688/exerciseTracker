package controllers

import (
	"io"

	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	*defaultController
}

func (ct *exerciseController) Create(c *gin.Context) {
	byt, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := ct.service.Create(byt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func (ct *exerciseController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "none",
	})
}
