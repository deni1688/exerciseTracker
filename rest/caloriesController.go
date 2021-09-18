package rest

import (
	"github.com/gin-gonic/gin"
)

type caloriesController struct {
	*defaultController
}

func (ct *caloriesController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreateCalories not implmented",
	})
}

func (ct *caloriesController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAllCalories not implmented",
	})
}
