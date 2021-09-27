package domain

func NewExerciseService(r ExerciseRepository) ExerciseService {
	return &exercisesService{r}
}

type exercisesService struct {
	repo ExerciseRepository
}

func (s *exercisesService) ListExercises() (*[]Exercise, error) {
	return s.repo.FindAll()
}

func (s *exercisesService) SaveExercise(ex *Exercise) (string, error) {
	newEx := NewExercise(ex.Category, ex.Name, ex.Weight, ex.Unit, ex.Sets, ex.Value)
	if err := newEx.Validate(); err != nil {
		return "", err
	}

	return s.repo.Create(newEx)
}
