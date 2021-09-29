package storage

import (
	"errors"
	"github.com/deni1688/exercise_tracker/domain"
	"github.com/deni1688/exercise_tracker/storage/jsondb"
	"github.com/deni1688/exercise_tracker/storage/sqlite"
)

func NewExerciseRepository(driver string) (domain.ExerciseRepository, error) {
	if driver == "sqlite" {
		return sqlite.NewExerciseRepository()
	}

	if driver == "json" {
		return jsondb.NewExerciseRepository()
	}

	return nil, errors.New("storage not supported")
}
