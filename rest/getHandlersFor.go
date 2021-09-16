package rest

import "deni1688/myHealthTrack/domain"

func GetHandlersFor(handlers string, service domain.Service) Handlers {
	dh := &defaultHandlers{service}

	switch handlers {
	case "exercises":
		return &exerciseHandlers{dh}
	case "weight":
		return &weightHandlers{dh}
	case "nutrition":
		return &nutritionHandlers{dh}
	default:
		return dh
	}
}
