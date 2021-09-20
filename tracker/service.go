package tracker

import (
	"encoding/json"
	"errors"
)

type Service interface {
	FindAll(category string, query string) ([]TrackEntry, error)
	FindOne(category string, id string) []byte
	Create(data []byte) (string, error)
	UpdateOne(category string, id string, data []byte) bool
}

func NewTrackerService(r Repository) Service {
	return &trackService{r}
}

type trackService struct {
	r Repository
}

func (s *trackService) FindAll(category string, query string) ([]TrackEntry, error) {
	exercises, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	return exercises, nil
}

func (s *trackService) FindOne(category string, id string) []byte {
	return nil
}

func (s *trackService) Create(data []byte) (string, error) {
	var te *TrackEntry
	if err := json.Unmarshal(data, &te); err != nil {
		return "", errors.New("Error trying to parse exercise data: " + err.Error())
	}

	return s.r.Create(NewTrackEntry(te.Category, te.Description, te.Value))
}

func (s *trackService) UpdateOne(category string, id string, data []byte) bool {
	return false
}
