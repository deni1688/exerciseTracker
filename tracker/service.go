package tracker

import (
	"encoding/json"
)

type Service interface {
	FindAll(category string, query string) []byte
	FindOne(category string, id string) []byte
	Create(category string, data []byte) string
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

func (s *trackService) Create(category string, data []byte) string {
	var te *trackEntry

	_ = json.Unmarshal(data, &te)

	return s.r.Create(te)
}

func (s *trackService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
