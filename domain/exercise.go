package domain

import (
	"github.com/go-playground/validator/v10"
	"time"
)

const (
	ExerciseCollection = "exercises"
)

var validate = validator.New()

type Exercise struct {
	ID        string    `json:"id,omitempty"`
	Category  string    `json:"category" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Sets      int32     `json:"sets"`
	Weight    float32   `json:"weight" validate:"gte=25"`
	Unit      string    `json:"unit" validate:"required"`
	Value     float32   `json:"value" validate:"required"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

func NewExercise(
	category, name string,
	weight float32,
	unit string,
	sets int32,
	value float32,
) (*Exercise, error) {
	ex := new(Exercise)
	ex.Category = category
	ex.Name = name
	ex.Weight = weight
	ex.Unit = unit
	ex.Sets = sets
	ex.Value = value

	err := validate.Struct(ex)
	if err != nil {
		return nil, err
	}

	return ex, nil
}
