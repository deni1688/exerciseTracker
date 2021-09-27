package domain

import (
	"errors"
	"strings"
	"time"
)

const ExerciseCollection = "exercises"
var ValidCategories = []string{"cardio", "calisthenics"}

type Exercise struct {
	ID        string    `json:"id,omitempty"`
	Category  string    `json:"category" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Weight    float32   `json:"weight" validate:"gte=25"`
	Sets      int32     `json:"sets"`
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
) *Exercise {
	ex := new(Exercise)
	ex.Category = category
	ex.Name = name
	ex.Weight = weight
	ex.Unit = unit
	ex.Sets = sets
	ex.Value = value

	return ex
}

func (ex *Exercise) Validate() error {

	if ex.Category == "" {
		return errors.New("category is required")
	}

	if !strings.Contains(strings.Join(ValidCategories, ""), ex.Category) {
		return errors.New("invalid category")
	}

	return nil
}
