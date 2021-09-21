package tracker

type ExerciseAggregate struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Duration int     `json:"duration,omitempty"`
	Distance int     `json:"distance,omitempty"`
	Reps     int     `json:"reps,omitempty"`
	Sets     int     `json:"sets,omitempty"`
}
