package domain

const Collection = "exercises"

type Exercise struct {
	ID       string  `json:"id" header:"id"`
	Category string  `json:"category" header:"category"`
	Name     string  `json:"name" header:"name"`
	Weight   float64 `json:"weight" header:"weight"`
	Duration int     `json:"duration,omitempty" header:"duration"`
	Distance int     `json:"distance,omitempty" header:"distance"`
	Reps     int     `json:"reps,omitempty" header:"reps"`
	Sets     int     `json:"sets,omitempty" header:"sets"`
	Created  int64   `json:"created" header:"created"`
}

func NewExercise(category, name string, weight float64) *Exercise {
	ex := new(Exercise)
	ex.Category = category
	ex.Name = name
	ex.Weight = weight
	return ex
}

func (ex *Exercise) ForCardio(duration, distance int) *Exercise {
	ex.Duration = duration
	ex.Distance = distance
	return ex
}

func (ex *Exercise) ForCalisthenics(reps, sets int) *Exercise {
	ex.Reps = reps
	ex.Sets = sets
	return ex
}
