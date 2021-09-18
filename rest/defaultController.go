package rest

import (
	"deni1688/myHealthTrack/tracker"

	"github.com/gin-gonic/gin"
)

type defaultController struct {
	service tracker.Service
	entity  string
}

func (ct *defaultController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create not implmented",
	})
}

func (ct *defaultController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAll not implmented",
	})
}

func (ct *defaultController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetOne not implmented for /" + ct.entity + "/" + c.Params.ByName("id"),
	})
}

func (ct *defaultController) UpdateOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateOne not implmented for /" + ct.entity + "/" + c.Params.ByName("id"),
	})
}
