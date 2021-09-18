package rest

import (
	"deni1688/myHealthTrack/tracker"

	"github.com/gin-gonic/gin"
)

type defaultController struct {
	service tracker.Service
	entity  string
}

func (ct *defaultController) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreate not implmented",
	})
}

func (ct *defaultController) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAll not implmented",
	})
}

func (ct *defaultController) HandleGetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetOne not implmented for /" + ct.entity + "/" + c.Params.ByName("id"),
	})
}

func (ct *defaultController) HandleUpdateOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleUpdateOne not implmented for /" + ct.entity + "/" + c.Params.ByName("id"),
	})
}
