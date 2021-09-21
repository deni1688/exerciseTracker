package exercises

type Repository interface {
	FindAll() ([]Entity, error)
	FindOne(id string) interface{}
	Create(te Entity) (string, error)
	UpdateOne(id string, object interface{}) bool
}
