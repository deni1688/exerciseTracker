package domain

type ExerciseRepository interface {
	FindAll() (*[]Exercise, error)
	Create(*Exercise) (string, error)
}

type ExerciseService interface {
	ListExercises() (*[]Exercise,error)
	SaveExercise(*Exercise) (string, error)
}
