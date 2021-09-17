package rest

import "github.com/gin-gonic/gin"

type api struct {
	engine *gin.Engine
}

func New(engine *gin.Engine) *api {
	return &api{engine}
}

func (a *api) Run(addr string) error {
	return a.engine.Run(addr)
}

func (a *api) InitExerciseRoutes(h Controller) *api {
	basePath := "/exercises"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST(basePath, h.HandleCreate)
	a.engine.GET(basePath+"/:id", h.HandleGetOne)
	return a
}

func (a *api) InitWeightRoutes(h Controller) *api {
	basePath := "/weight"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST(basePath, h.HandleCreate)
	a.engine.GET(basePath+"/:id", h.HandleGetOne)
	return a
}

func (a *api) InitCaloriesRoutes(h Controller) *api {
	basePath := "/calories"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST(basePath, h.HandleCreate)
	a.engine.GET(basePath+"/:id", h.HandleGetOne)
	return a
}
