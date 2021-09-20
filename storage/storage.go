package storage

import (
	"deni1688/myHealthTrack/storage/jsondb"
	"deni1688/myHealthTrack/tracker"
	"errors"
	"os"
)

func GetRepository(storage string) (tracker.Repository, error) {
	if storage == "jsondb" {
		jsonDBDir := os.Getenv("JSON_DB_DIR")

		if jsonDBDir == "" {
			return nil, errors.New("Error getting the JSON_DB_DIR env variable!")
		}

		return jsondb.New(jsonDBDir), nil
	}

	return nil, errors.New(storage + " is not supported!")
}
