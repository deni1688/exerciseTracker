package domain

import "time"

const (
	ExerciseCollection   = "exercises"
	CardioCategory       = "cardio"
	CalisthenicsCategory = "calisthenics"
)

type Exercise struct {
	ID             string    `json:"id,omitempty"`
	ScheduleID     string    `json:"schedule_id"`
	Category       string    `json:"category,omitempty"`
	Name           string    `json:"name,omitempty"`
	Weight         float32   `json:"weight,omitempty"`
	MinuteDuration int32     `json:"minute_duration,omitempty"`
	KmDistance     float32   `json:"km_distance,omitempty"`
	Reps           int32     `json:"reps,omitempty"`
	Sets           int32     `json:"sets,omitempty"`
	Created        time.Time `json:"created,omitempty"`
}

func NewExercise(category, name string, weight float32) *Exercise {
	ex := new(Exercise)
	ex.Category = category
	ex.Name = name
	ex.Weight = weight
	return ex
}

func (ex *Exercise) ForCardio(duration int32, distance float32) *Exercise {
	ex.MinuteDuration = duration
	ex.KmDistance = distance
	return ex
}

func (ex *Exercise) ForCalisthenics(reps, sets int32) *Exercise {
	ex.Reps = reps
	ex.Sets = sets
	return ex
}
