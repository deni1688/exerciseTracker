package tracker

type Service interface {
	FindAll(entity string, query string) []interface{}
	FindOne(entity string, id string) interface{}
	Create(entity string, object interface{}) bool
	UpdateOne(entity string, id string, object interface{}) bool
}

type Repository interface {
	FindAll(query string) []interface{}
	FindOne(id string) []interface{}
	Create(object interface{}) bool
	UpdateOne(entity string, id string, object interface{}) bool
}
