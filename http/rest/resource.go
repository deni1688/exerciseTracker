package rest

import "github.com/gin-gonic/gin"

type Resource interface {
	With(ct Controller)
}

type resourceFactory struct {
	router     *gin.Engine
	collection string
}

func NewResource(router *gin.Engine, category string) Resource {
	return &resourceFactory{router, category}
}

func (r *resourceFactory) With(ct Controller) {
	r.router.GET(resourcePathFromCollection(r.collection), ct.GetAll)
	r.router.POST(resourcePathFromCollection(r.collection), ct.Create)
	r.router.GET(resourcePathFromCollection(r.collection, ":id"), ct.GetOne)
	r.router.PUT(resourcePathFromCollection(r.collection, ":id"), ct.UpdateOne)
}

func resourcePathFromCollection(category string, params ...string) string {
	resource := "/" + category

	for _, p := range params {
		resource += "/" + p
	}

	return resource
}
