package storage

import (
	"deni1688/exerciseTracker/exercise"
	"deni1688/exerciseTracker/storage/jsondb"
	"deni1688/exerciseTracker/storage/sqlite"
	"errors"
)

func NewExerciseRepository(driver string) (exercise.Repository, error) {
	if driver == "sqlite" {
		return sqlite.NewExerciseRepository()
	}

	if driver == "json" {
		return jsondb.NewExerciseRepository()
	}

	return nil, errors.New("storage not supported")
}
