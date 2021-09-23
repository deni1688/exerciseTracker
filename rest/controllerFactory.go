package rest

import (
	"deni1688/exerciseTracker/exercise"
)

type ControllerFactory interface {
	For(resource string) Controller
}

type controllerFactory struct {
	service exercise.Service
}

func NewControllerFactory(service exercise.Service) ControllerFactory {
	return &controllerFactory{service}
}

func (cf controllerFactory) For(resource string) Controller {
	dc := &defaultController{cf.service, resource}
	if resource == exercise.Collection {
		return &exerciseController{dc}
	}
	return dc
}
