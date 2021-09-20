package jsondb

import (
	"deni1688/myHealthTrack/tracker"
	"fmt"

	uuid "github.com/google/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
)

type jsondb struct {
	db *scribble.Driver
}

func New(dir string) tracker.Repository {
	db, err := scribble.New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	return &jsondb{db}
}

func (jdb *jsondb) FindAll(query string) []interface{} {
	return nil
}

func (jdb *jsondb) FindOne(id string) interface{} {
	return nil
}

func (jdb *jsondb) Create(object interface{}) string {
	id := uuid.New().String()

	if err := jdb.db.Write("trackpoints", id, object); err != nil {
		fmt.Println("Error", err)
	}

	return id
}

func (jdb *jsondb) UpdateOne(id string, object interface{}) bool {
	return false
}
