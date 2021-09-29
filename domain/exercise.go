package domain

import (
	"errors"
	"strings"
	"time"
)

const ExerciseCollection = "exercises"

var (
	ValidCategories     = []string{"cardio", "calisthenics"}
	ErrCategoryRequired = errors.New("category is required")
	ErrInvalidCategory  = errors.New("invalid category")
	ErrInvalidWeight    = errors.New("invalid weight")
	ErrInvalidUnit      = errors.New("invalid unit")
	ErrInvalidSets      = errors.New("invalid sets")
	ErrInvalidValue     = errors.New("invalid value")
)

type Exercise struct {
	ID        string    `json:"id,omitempty"`
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	Weight    float32   `json:"weight"`
	Sets      int32     `json:"sets"`
	Unit      string    `json:"unit"`
	Value     float32   `json:"value"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

type ExerciseRequest struct {
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	Weight    float32   `json:"weight"`
	Sets      int32     `json:"sets"`
	Unit      string    `json:"unit"`
	Value     float32   `json:"value"`
}

func newExercise(
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
		return ErrCategoryRequired
	}
	if !strings.Contains(strings.Join(ValidCategories, ""), ex.Category) {
		return ErrInvalidCategory
	}
	if ex.Weight <= 0 {
		return ErrInvalidWeight
	}
	if ex.Unit == "" {
		return ErrInvalidUnit
	}
	if ex.Sets < 0 {
		return ErrInvalidSets
	}
	if ex.Value <= 0 {
		return ErrInvalidValue
	}

	return nil
}
