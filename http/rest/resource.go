package rest

import "github.com/gin-gonic/gin"

type Resource interface {
	With(ct Controller)
}

type resourceFactory struct {
	router *gin.Engine
	category string
}

func NewResource(router *gin.Engine, category string) Resource {
	return &resourceFactory{router, category}
}

func (r *resourceFactory) With(ct Controller) {
	r.router.GET(resourcePathFromCategory(r.category), ct.GetAll)
	r.router.POST(resourcePathFromCategory(r.category), ct.Create)
	r.router.GET(resourcePathFromCategory(r.category, ":id"), ct.GetOne)
	r.router.PUT(resourcePathFromCategory(r.category, ":id"), ct.UpdateOne)
}

func resourcePathFromCategory(category string, params ...string) string {
	resource := "/" + category

	for _, p := range params {
		resource += "/" + p
	}

	return resource
}
