package http

import (
	"deni1688/exercise_tracker/domain"
)

type HandlerFactory interface {
	For(resource string) Handler
}

type handlerFactory struct {
	service domain.Service
}

func NewHandlerFactory(service domain.Service) HandlerFactory {
	return &handlerFactory{service}
}

func (cf handlerFactory) For(resource string) Handler {
	dc := &defaultHandler{cf.service, resource}
	if resource == domain.Collection {
		return &exerciseHandler{dc}
	}
	return dc
}
