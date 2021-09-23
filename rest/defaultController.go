package rest

import (
	"deni1688/exerciseTracker/domain"

	"github.com/gin-gonic/gin"
)

type defaultController struct {
	service  domain.Service
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
