package rest

import "deni1688/myHealthTrack/tracker"

type ControllerFactory interface {
    Create(resource string) Controller
}

type controllerFactory struct {
    service tracker.Service
}

func NewControllerFactory(service tracker.Service) ControllerFactory {
   return &controllerFactory{service}
}

func (cf controllerFactory) Create(resource string) Controller {
	dc := &defaultController{cf.service, resource}

	return &exerciseController{dc}
}
