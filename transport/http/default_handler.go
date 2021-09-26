package http

import (
	"deni1688/exercise_tracker/domain"

	"github.com/gin-gonic/gin"
)

type defaultHandler struct {
	service  domain.Service
	resource string
}

func (ct *defaultHandler) GetResource() string {
	return ct.resource
}

func (ct *defaultHandler) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "With not implemented for " + ct.resource,
	})
}

func (ct *defaultHandler) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAll not implemented " + ct.resource,
	})
}
