package tracker

type Repository interface {
	FindAll(query string) []interface{}
	FindOne(id string) []interface{}
	Create(object interface{}) bool
	UpdateOne(category string, id string, object interface{}) bool
}
