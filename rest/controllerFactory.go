package rest

import "deni1688/myHealthTrack/tracker"

type ControllerFactory interface {
    Create(entity string) Controller
}

type controllerFactory struct {
    service tracker.Service
}

func NewControllerFactory(service tracker.Service) ControllerFactory{
   return &controllerFactory{service}
}

func (cf controllerFactory) Create(entity string) Controller {
	dc := &defaultController{cf.service, entity}

    if entity == tracker.EXERCISE {
        return &exerciseController{dc}
    }
    if entity == tracker.WEIGHT {
        return &weightController{dc}
    }
    if entity == tracker.CALORIES {
        return &caloriesController{dc}
    }

	return dc
}
