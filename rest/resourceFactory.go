package rest

import "github.com/gin-gonic/gin"

type ResourceFactory interface {
	Create(entity string, ct Controller)
}

type resourceFactory struct {
	router *gin.Engine
}

func NewResourceFactory(router *gin.Engine) ResourceFactory {
	return &resourceFactory{router}
}

func (rf *resourceFactory) Create(entity string, ct Controller) {
	rf.router.GET(resourceFromEntity(entity), ct.GetAll)
	rf.router.POST(resourceFromEntity(entity), ct.Create)
	rf.router.GET(resourceFromEntity(entity, ":id"), ct.GetOne)
	rf.router.PUT(resourceFromEntity(entity, ":id"), ct.UpdateOne)
}

func resourceFromEntity(entity string, params ...string) string {
	resource := "/" + entity

	for _, p := range params {
		resource += "/" + p
	}

	return resource
}
