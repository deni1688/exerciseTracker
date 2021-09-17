package rest

import (
	"deni1688/myHealthTrack/tracker"

	"github.com/gin-gonic/gin"
)

type defaultController struct {
	service tracker.Service
	entity  string
}

func (h *defaultController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreate not implmented",
	})
}

func (h *defaultController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAll not implmented",
	})
}

func (h *defaultController) HandleGetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetOne not implmented",
	})
}
