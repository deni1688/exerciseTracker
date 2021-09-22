package domain

type Repository interface {
	FindAll() (*[]Exercise, error)
	Create(*Exercise) (string, error)
}

type Service interface {
	FindAll() (*[]Exercise, error)
	Create(data *Exercise) (string, error)
}
