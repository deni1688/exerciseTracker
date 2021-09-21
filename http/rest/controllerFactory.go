package rest

import "deni1688/myHealthTrack/exercises"

type ControllerFactory interface {
    Create(resource string) Controller
}

type controllerFactory struct {
    service exercises.Service
}

func NewControllerFactory(service exercises.Service) ControllerFactory {
   return &controllerFactory{service}
}

func (cf controllerFactory) Create(resource string) Controller {
	dc := &defaultController{cf.service, resource}

	return &exerciseController{dc}
}
