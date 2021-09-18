package rest

import (
	"github.com/gin-gonic/gin"
)

type weightController struct {
	*defaultController
}

func (ct *weightController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateWeight not implemented",
	})
}

func (ct *weightController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllWeight not implemented",
	})
}
