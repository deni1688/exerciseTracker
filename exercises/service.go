package exercises

import "errors"

var invalidCategoryErr = errors.New("invalid exercise category")

type Service interface {
	FindAll() (*[]Exercise, error)
	FindOne(category string, id string) []byte
	Create(data *Exercise) (string, error)
	UpdateOne(category string, id string, data []byte) bool
}

func NewExerciseService(r Repository) Service {
	return &exercisesService{r}
}

type exercisesService struct {
	repo Repository
}

func (s *exercisesService) FindAll() (*[]Exercise, error) {
	_, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *exercisesService) FindOne(category string, id string) []byte {
	return nil
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

func (s *exercisesService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
