package tracker

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

const Collection = "exercises"

type Exercise interface {
	GetID() string
	Encode() ([]byte, error)
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
func (ex *exercise) Encode() ([]byte, error) {
	return json.Marshal(ex)
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

