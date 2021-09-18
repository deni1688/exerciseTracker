package rest

import (
	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	*defaultController
}

func (ct *exerciseController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateExercise not implemented",
	})
}

func (ct *exerciseController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllExercises not implemented",
	})
}
