package jsondb

import (
	"deni1688/myHealthTrack/tracker"
	"fmt"

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
	return ""
}

func (jdb *jsondb) UpdateOne(id string, object interface{}) bool {
	return false
}
