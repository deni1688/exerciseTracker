package exercises

import (
	"github.com/google/uuid"
	"time"
)

const Collection = "exercises"

type Exercise struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Duration int     `json:"duration,omitempty"`
	Distance int     `json:"distance,omitempty"`
	Reps     int     `json:"reps,omitempty"`
	Sets     int     `json:"sets,omitempty"`
	Created  int64   `json:"created"`
}

func NewExercise(category, name string, weight float64) *Exercise {
	ex := new(Exercise)
	ex.ID = uuid.New().String()
	ex.Created = time.Now().Unix()
	ex.Category = category
	ex.Name = name
	ex.Weight = weight
	return ex
}

func (ex *Exercise) ForCardio(duration, distance int) *Exercise {
	ex.Duration = duration
	ex.Distance = distance
	return ex
}

func (ex *Exercise) ForCalisthenics(reps, sets int) *Exercise {
	ex.Reps = reps
	ex.Sets = sets
	return ex
}
