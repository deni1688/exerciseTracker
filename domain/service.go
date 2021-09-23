package domain

import "errors"

var invalidCategoryErr = errors.New("invalid exercise category")

func NewExerciseService(r Repository) Service {
	return &exercisesService{r}
}

type exercisesService struct {
	repo Repository
}

func (s *exercisesService) ListExercises() (*[]Exercise, error) {
	return s.repo.FindAll()
}

func (s *exercisesService) SaveExercise(ex *Exercise) (string, error) {
	if ex.Category == "cardio" {
		return s.repo.Create(NewExercise(
			ex.Category,
			ex.Name,
			ex.Weight,
		).ForCardio(ex.Duration, ex.Distance))
	}

	if ex.Category == "calisthenics" {
		return s.repo.Create(NewExercise(
			ex.Category,
			ex.Name,
			ex.Weight,
		).ForCalisthenics(ex.Reps, ex.Sets))
	}

	return "", invalidCategoryErr
}
