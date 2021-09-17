package rest

import "deni1688/myHealthTrack/tracker"

func CreateControllerFor(entity string, service tracker.Service) Controller {
	dh := &defaultController{service, entity}

	controller, found := map[string]Controller{
		"exercise":  &exerciseController{dh},
		"weight":    &weightController{dh},
		"calories": &caloriesController{dh},
	}[entity]

	if !found {
		return dh
	}

	return controller
}
