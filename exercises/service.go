package exercises

import (
	"encoding/json"
	"errors"
)

type Service interface {
	FindAll() ([]Entity, error)
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

func (s *exercisesService) FindAll() ([]Entity, error) {
	exercises, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func (s *exercisesService) FindOne(category string, id string) []byte {
	return nil
}

func (s *exercisesService) Create(data []byte) (string, error) {
	var ex *Aggregate

	if err := json.Unmarshal(data, &ex); err != nil {
		return "", errors.New("Error trying to parse exercise data: " + err.Error())
	}

	return s.repo.Create(newExercise(ex))
}

func (s *exercisesService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
