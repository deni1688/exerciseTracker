package tracker

const (
	EXERCISE = "exercise"
	WEIGHT   = "weight"
)

type trackEntry struct {
	category string
	value    float64
	description string
}

func NewTrackEntry(category, description string, value float64) *trackEntry {
	return &trackEntry{category,value,description}
}
