package rest

import (
	"github.com/gin-gonic/gin"
)

type weightController struct {
	*defaultController
}

func (ct *weightController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateWeight not implmented",
	})
}

func (ct *weightController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllWeight not implmented",
	})
}
