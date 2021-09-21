package tracker

type ExerciseAggregate struct {
	*exercise
	*cardio
	*calisthenics
}
