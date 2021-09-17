package rest

import "deni1688/myHealthTrack/tracker"

func CreateHandlersFor(entity string, service tracker.Service) Handlers {
	dh := &defaultHandlers{service, entity}

	handlers, found := map[string]Handlers{
		"exercise":  &exerciseHandlers{dh},
		"weight":    &weightHandlers{dh},
		"nutrition": &nutritionHandlers{dh},
	}[entity]

	if !found {
		return dh
	}

	return handlers
}
