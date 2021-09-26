package domain

import "time"

type Schedule struct {
	ID string
	Exercises []Exercise
	Created time.Time
}
