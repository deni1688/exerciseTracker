package rest

import (
	"github.com/deni1688/exercise_tracker/domain"
)

func GetHandlerFor(service domain.ExerciseService, resource string) Handler {
	if resource == domain.ExerciseCollection {
		return &exerciseHandler{service, resource}
	}
	return nil
}
