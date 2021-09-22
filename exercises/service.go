package exercises

import (
	"encoding/json"
	"errors"
)

type Service interface {
	FindAll() (*[]Exercise, error)
	FindOne(category string, id string) []byte
	Create(data []byte) (string, error)
	UpdateOne(category string, id string, data []byte) bool
}

func NewService(r Repository) Service {
	return &exercisesService{r}
}

type exercisesService struct {
	repo Repository
}

func (s *exercisesService) FindAll() (*[]Exercise, error) {
	exercises, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func (s *exercisesService) FindOne(category string, id string) []byte {
	return nil
}

func (s *exercisesService) Create(er *Exercise) (string, error) {
	ex := new(Exercise)

	if er.Category == "cardio" {
		ex = NewExercise(er.Category, er.Name, er.Weight).ForCardio(er.Duration, er.Distance)
	}

	if er.Category == "calisthenics" {
		ex = NewExercise(er.Category, er.Name, er.Weight).ForCalisthenics(er.Reps, er.Sets)
	}

	return s.repo.Create(ex)
}

func (s *exercisesService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
