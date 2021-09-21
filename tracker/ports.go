package tracker

type Repository interface {
	FindAll() ([]Exercise, error)
	FindOne(id string) interface{}
	Create(te Exercise) (string, error)
	UpdateOne(id string, object interface{}) bool
}
