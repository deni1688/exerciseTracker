package tracker

import (
	"github.com/google/uuid"
	"time"
)

const (
	EXERCISE = "exercise"
	WEIGHT   = "weight"
)

type TrackEntry struct {
	ID          string  `json:"id"`
	Category    string  `json:"category"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
	Created     string  `json:"created"`
}

func NewTrackEntry(category, description string, value float64) *TrackEntry {
	id := uuid.New().String()
	created := time.Now().Format(time.RFC3339Nano)
	return &TrackEntry{id, category, value, description, created}
}
