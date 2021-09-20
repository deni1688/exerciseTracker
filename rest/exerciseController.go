package rest

import (
	"io"

	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	*defaultController
}

func (ct *exerciseController) Create(c *gin.Context) {
	byt, _ := io.ReadAll(c.Request.Body)

	c.JSON(200, gin.H{
		"id": ct.service.Create(ct.category, byt),
	})
}

func (ct *exerciseController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllExercises not implemented",
	})
}
