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
		"message": "Create not implemented for " + ct.entity,
	})
}

func (ct *defaultController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAll not implemented " + ct.entity,
	})
}

func (ct *defaultController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetOne not implemented for /" + ct.entity + "/" + c.Params.ByName("id"),
	})
}

func (ct *defaultController) UpdateOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateOne not implemented for /" + ct.entity + "/" + c.Params.ByName("id"),
	})
}
