package tracker

type Repository interface {
	FindAll(query string) []interface{}
	FindOne(id string) interface{}
	Create(object interface{}) string
	UpdateOne(id string, object interface{}) bool
}
