package rest

import "deni1688/myHealthTrack/domain"

func CreateHandlersFor(entity string, service domain.Service) Handlers {
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
