package rest

import "github.com/gin-gonic/gin"

type api struct {
	engine *gin.Engine
}

func New(engine *gin.Engine) *api {
	return &api{engine}
}

func (a *api) Run() {
	a.engine.Run()
}

func (a *api) InitExerciseRoutes(h Handlers) *api {
	basePath := "/exercises"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST("/exercises", h.HandleCreate)
	a.engine.GET("/exercises/:id", h.HandleGetOne)
	return a
}

func (a *api) InitWeightRoutes(h Handlers) *api {
	basePath := "/weight"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST("/weight", h.HandleCreate)
	return a
}

func (a *api) InitNutritionRoutes(h Handlers) *api {
	basePath := "/nutrition"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST(basePath, h.HandleCreate)
	return a
}
