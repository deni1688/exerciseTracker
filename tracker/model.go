package tracker

import (
	"github.com/google/uuid"
	"time"
)

const Collection = "exercises"

type Exercise interface {
	GetID() string
}

type exercise struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Category string  `json:"category"`
	Created  int64   `json:"created"`
}

type cardio struct {
	*exercise
	Duration int `json:"duration"`
	Distance int `json:"distance"`
}

type calisthenics struct {
	*exercise
	Reps int `json:"reps"`
	Sets int `json:"sets"`
}

func (ex *exercise) GetID() string {
	return ex.ID
}

func newExercise(er *ExerciseAggregate) Exercise {
	id := uuid.New().String()
	created := time.Now().Unix()

	ex := &exercise{id, er.Name, er.Weight, er.Category, created}

	if er.Category == "cardio" {
		return &cardio{ex, er.Duration, er.Distance}
	}

	if er.Category == "calisthenics" {
		return &calisthenics{ex, er.Reps, er.Sets}
	}

	return ex
}

type ExerciseAggregate struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Duration int     `json:"duration,omitempty"`
	Distance int     `json:"distance,omitempty"`
	Reps     int     `json:"reps,omitempty"`
	Sets     int     `json:"sets,omitempty"`
}
