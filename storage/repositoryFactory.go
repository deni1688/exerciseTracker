package storage

import (
	"deni1688/exerciseTracker/domain"
	"deni1688/exerciseTracker/storage/jsondb"
	"deni1688/exerciseTracker/storage/sqlite"
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
