package rest

import "deni1688/myHealthTrack/tracker"

type ControllerFactory interface {
    Create(category string) Controller
}

type controllerFactory struct {
    service tracker.Service
}

func NewControllerFactory(service tracker.Service) ControllerFactory{
   return &controllerFactory{service}
}

func (cf controllerFactory) Create(category string) Controller {
	dc := &defaultController{cf.service, category}

    if category == tracker.EXERCISE {
        return &exerciseController{dc}
    }

    if category == tracker.WEIGHT {
        return &weightController{dc}
    }


	return dc
}
