package tracker

const (
	EXERCISE = "exercise"
	WEIGHT   = "weight"
)

type trackEntry struct {
	ID          string  `json:"id"`
	Category    string  `json:"category"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
}

func NewTrackEntry(category, description string, value float64) *trackEntry {
	return &trackEntry{"", category, value, description}
}
