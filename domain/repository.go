package domain

type Repository interface {
	FindAll(query string) []interface{}
	FindOne(id string) []interface{}
	Create(object interface{}) bool
}
