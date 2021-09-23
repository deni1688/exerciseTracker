package domain

type Repository interface {
	FindAll() (*[]Exercise, error)
	Create(*Exercise) (string, error)
}

type Service interface {
	ListExercises() (*[]Exercise,error)
	SaveExercise(*Exercise) (string, error)
}
