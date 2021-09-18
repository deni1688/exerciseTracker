package rest

import (
	"github.com/gin-gonic/gin"
)

type caloriesController struct {
	*defaultController
}

func (ct *caloriesController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateCalories not implmented",
	})
}

func (ct *caloriesController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllCalories not implmented",
	})
}
