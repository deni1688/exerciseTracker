package rest

import (
	"deni1688/exercise_tracker/domain"
)

type ControllerFactory interface {
	For(resource string) Controller
}

type controllerFactory struct {
	service domain.Service
}

func NewControllerFactory(service domain.Service) ControllerFactory {
	return &controllerFactory{service}
}

func (cf controllerFactory) For(resource string) Controller {
	dc := &defaultController{cf.service, resource}
	if resource == domain.Collection {
		return &exerciseController{dc}
	}
	return dc
}
