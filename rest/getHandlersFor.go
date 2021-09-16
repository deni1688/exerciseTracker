package rest

import "deni1688/myHealthTrack/domain"

func GetHandlersFor(entity string, service domain.Service) Handlers {
	dh := &defaultHandlers{service, entity}

	switch entity {
	case "exercise":
		return &exerciseHandlers{dh}
	case "weight":
		return &weightHandlers{dh}
	case "nutrition":
		return &nutritionHandlers{dh}
	default:
		return dh
	}
}
