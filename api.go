package main

import "github.com/gin-gonic/gin"

type api struct {
	engine *gin.Engine
}

func newApi(engine *gin.Engine) *api {
	return &api{engine}
}

func (a *api) Run() {
	a.engine.Run()
}

func (a *api) InitExerciseApi(h Handlers) *api {
	basePath := "/exercises"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST("/exercises", h.HandleCreate)
	return a
}

func (a *api) InitWeightApi(h Handlers) *api {
	basePath := "/weight"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST("/weight", h.HandleCreate)
	return a
}

func (a *api) InitNutritionApi(h Handlers) *api {
	basePath := "/nutrition"
	a.engine.GET(basePath, h.HandleGetAll)
	a.engine.POST(basePath, h.HandleCreate)
	return a
}
