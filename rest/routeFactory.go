package rest

import "github.com/gin-gonic/gin"

type RouteFactory interface {
	Create(entity string, ct Controller)
}

type routeFactory struct {
	router *gin.Engine
}

func NewRouteFactory(router *gin.Engine) RouteFactory {
	return &routeFactory{router}
}

func (rf *routeFactory) Create(entity string, ct Controller) {
	rf.router.GET(resourceFromEntity(entity), ct.HandleGetAll)
	rf.router.POST(resourceFromEntity(entity), ct.HandleCreate)
	rf.router.GET(resourceFromEntity(entity, ":id"), ct.HandleGetOne)
	rf.router.PUT(resourceFromEntity(entity, ":id"), ct.HandleUpdateOne)
}

func resourceFromEntity(entity string, params ...string) string {
	resource := "/" + entity

	for _, p := range params {
		resource += "/" + p
	}

	return resource
}
