package controllers

import (
	"deni1688/exerciseTracker/exercises"
)

type ControllerFactory interface {
    For(resource string) Controller
}

type controllerFactory struct {
    service exercises.Service
}

func NewControllerFactory(service exercises.Service) ControllerFactory {
   return &controllerFactory{service}
}

func (cf controllerFactory) For(resource string) Controller {
	dc := &defaultController{cf.service, resource}
	if resource == exercises.Collection {
		return &exerciseController{dc}
	}
	return dc
}
