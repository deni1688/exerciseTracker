package rest

import (
	"github.com/gin-gonic/gin"
)

func NewResource(router *gin.Engine, handler Handler) {
	resource := handler.GetResource()
	router.GET(getResourcePath(resource), handler.GetAll)
	router.POST(getResourcePath(resource), handler.Create)
}

func getResourcePath(category string, params ...string) string {
	resource := "/" + category
	for _, p := range params {
		resource += "/" + p
	}
	return resource
}
