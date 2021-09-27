package domain

import "log"

func NewExerciseService(r ExerciseRepository, b ExerciseBroker) ExerciseService {
	return &exercisesService{r, b}
}

type exercisesService struct {
	repo   ExerciseRepository
	broker ExerciseBroker
}

func (s *exercisesService) ListExercises() (*[]Exercise, error) {
	return s.repo.FindAll()
}

func (s *exercisesService) SaveExercise(ex *Exercise) (string, error) {
	newEx := NewExercise(ex.Category, ex.Name, ex.Weight, ex.Unit, ex.Sets, ex.Value)
	if err := newEx.Validate(); err != nil {
		return "", err
	}

	id, err := s.repo.Create(newEx)
	if err != nil {
		return "", err
	}

	if err := s.broker.Publish(newEx); err != nil {
		log.Println("error publishing new exercise", err)
	}

	return id, nil
}
