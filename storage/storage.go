package storage

import (
	"deni1688/myHealthTrack/storage/jsondb"
	"deni1688/myHealthTrack/tracker"
	"errors"
	"os"
)

func NewRepository() (tracker.Repository, error) {
	storage := os.Getenv("STORAGE")

	if storage == "jsonDB" {
		jsonDBDir := os.Getenv("JSON_DB_DIR")

		if jsonDBDir == "" {
			return nil, errors.New("error getting the JSON_DB_DIR env variable")
		}

		return jsondb.New(jsonDBDir), nil
	}

	return nil, errors.New(storage + " is not supported!")
}
