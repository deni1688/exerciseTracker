package tracker

import (
	"encoding/json"
	"errors"
)

type Service interface {
	FindAll() ([]Exercise, error)
	FindOne(category string, id string) []byte
	Create(data []byte) (string, error)
	UpdateOne(category string, id string, data []byte) bool
}

func NewTrackerService(r Repository) Service {
	return &trackService{r}
}

type trackService struct {
	repo Repository
}

func (s *trackService) FindAll() ([]Exercise, error) {
	exercises, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func (s *trackService) FindOne(category string, id string) []byte {
	return nil
}

func (s *trackService) Create(data []byte) (string, error) {
	var ex *ExerciseAggregate

	if err := json.Unmarshal(data, &ex); err != nil {
		return "", errors.New("Error trying to parse exercise data: " + err.Error())
	}

	return s.repo.Create(newExercise(ex))
}

func (s *trackService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
