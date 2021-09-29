package domain

type ExerciseRepository interface {
	FindAll() (*[]Exercise, error)
	Create(*Exercise) (string, error)
}

type ExerciseService interface {
	ListExercises() (*[]Exercise, error)
	SaveExercise(request *ExerciseRequest) (string, error)
}

type ExerciseBroker interface {
	PublishCreated(*Exercise) error
	PublishUpdated(*Exercise) error
}
