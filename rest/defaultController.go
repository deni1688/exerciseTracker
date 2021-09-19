package rest

import (
	"deni1688/myHealthTrack/tracker"

	"github.com/gin-gonic/gin"
)

type defaultController struct {
	service tracker.Service
	category  string
}

func (ct *defaultController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "With not implemented for " + ct.category,
	})
}

func (ct *defaultController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAll not implemented " + ct.category,
	})
}

func (ct *defaultController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetOne not implemented for /" + ct.category + "/" + c.Params.ByName("id"),
	})
}

func (ct *defaultController) UpdateOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateOne not implemented for /" + ct.category + "/" + c.Params.ByName("id"),
	})
}
