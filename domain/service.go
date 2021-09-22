package domain

import "errors"

var invalidCategoryErr = errors.New("invalid exercise category")

func NewExerciseService(r Repository) Service {
	return &exercisesService{r}
}

type exercisesService struct {
	repo Repository
}

func (s *exercisesService) FindAll() (*[]Exercise, error) {
	return s.repo.FindAll()
}

func (s *exercisesService) Create(er *Exercise) (string, error) {
	if er.Category == "cardio" {
		return s.repo.Create(NewExercise(er.Category, er.Name, er.Weight).ForCardio(er.Duration, er.Distance))
	}

	if er.Category == "calisthenics" {
		return s.repo.Create(NewExercise(er.Category, er.Name, er.Weight).ForCalisthenics(er.Reps, er.Sets))
	}

	return "", invalidCategoryErr
}
