package rest

import "github.com/gin-gonic/gin"

type Controller interface {
	HandleCreate(c *gin.Context)
	HandleGetAll(c *gin.Context)
	HandleGetOne(c *gin.Context)
}
