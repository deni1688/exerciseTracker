package http

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetResource() string
}
