package storage

import (
	"deni1688/exercise_tracker/domain"
	"deni1688/exercise_tracker/storage/jsondb"
	"deni1688/exercise_tracker/storage/sqlite"
	"errors"
)

func NewExerciseRepository(driver string) (domain.Repository, error) {
	if driver == "sqlite" {
		return sqlite.NewExerciseRepository()
	}

	if driver == "json" {
		return jsondb.NewExerciseRepository()
	}

	return nil, errors.New("storage not supported")
}
