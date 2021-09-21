package rest

import (
	"deni1688/exerciseTracker/exercises"

	"github.com/gin-gonic/gin"
)

type defaultController struct {
	service  exercises.Service
	resource string
}

func (ct *defaultController) GetResource() string {
	return ct.resource
}

func (ct *defaultController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "With not implemented for " + ct.resource,
	})
}

func (ct *defaultController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAll not implemented " + ct.resource,
	})
}

func (ct *defaultController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetOne not implemented for /" + ct.resource + "/" + c.Params.ByName("id"),
	})
}

func (ct *defaultController) UpdateOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateOne not implemented for /" + ct.resource + "/" + c.Params.ByName("id"),
	})
}
