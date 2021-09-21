package rest

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	UpdateOne(c *gin.Context)
	GetResource() string
}
