package rest

import "github.com/gin-gonic/gin"

type Handler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetResource() string
}
