package exercises

type Repository interface {
	FindAll() (*[]Exercise, error)
	FindOne(string) (*Exercise, error)
	Create(*Exercise) (string, error)
	UpdateOne(string, *Exercise) (bool, error)
}
