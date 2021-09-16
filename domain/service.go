package domain

type Service interface {
	FindAll(entity string, query string) []interface{}
	FindOne(entity string, id string) interface{}
	Create(entity string, object interface{}) bool
}

func NewService() Service {
	return &myHealthTrackService{}
}

type myHealthTrackService struct {
	r *Repository
}

func (s *myHealthTrackService) FindAll(entity string, query string) []interface{} {
	return []interface{}{}
}

func (s *myHealthTrackService) FindOne(entity string, id string) interface{} {
	return nil
}

func (s *myHealthTrackService) Create(entity string, object interface{}) bool {
	return false
}
