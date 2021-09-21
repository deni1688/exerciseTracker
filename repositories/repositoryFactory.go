package repositories

import (
	"deni1688/exerciseTracker/exercises"
	"deni1688/exerciseTracker/repositories/jsondb"
	"errors"
	"os"
)

func NewRepository() (exercises.Repository, error) {
	storage := os.Getenv("STORAGE")

	if storage == "jsonDB" {
		jsonDBDir := os.Getenv("JSON_DB_DIR")

		if jsonDBDir == "" {
			return nil, errors.New("error getting the JSON_DB_DIR env variable")
		}

		return jsondb.New(jsonDBDir), nil
	}

	if storage != "" {
		storage = "Undefined"
	}

	return nil, errors.New(storage + " is not supported!")
}
