package tracker

type Service interface {
	FindAll(category string, query string) []byte
	FindOne(category string, id string) []byte
	Create(category string, data []byte) bool
	UpdateOne(category string, id string, data []byte) bool
}

func NewTrackerService(r Repository) Service {
	return &trackService{r}
}

type trackService struct {
	r Repository
}

func (s *trackService) FindAll(category string, query string) []byte {
	return nil
}

func (s *trackService) FindOne(category string, id string) []byte {
	return nil
}

func (s *trackService) Create(category string, data []byte) bool {
	return false
}

func (s *trackService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
