package rest

import (
	. "deni1688/myHealthTrack/http/rest/controllers"
	"github.com/gin-gonic/gin"
)


func NewResource(router *gin.Engine, controller Controller) {
	resource := controller.GetResource()
	router.GET(getResourcePath(resource), controller.GetAll)
	router.POST(getResourcePath(resource), controller.Create)
	router.GET(getResourcePath(resource, ":id"), controller.GetOne)
	router.PUT(getResourcePath(resource, ":id"), controller.UpdateOne)
}

func getResourcePath(category string, params ...string) string {
	resource := "/" + category
	for _, p := range params {
		resource += "/" + p
	}
	return resource
}
