package jsondb

import (
	"deni1688/myHealthTrack/tracker"
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
)

type jsonDB struct {
	db *scribble.Driver
}

func New(dir string) tracker.Repository {
	db, err := scribble.New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	return &jsonDB{db}
}

func (jdb *jsonDB) FindAll() ([]tracker.Exercise, error) {
	return nil, nil
}

func (jdb *jsonDB) FindOne(id string) interface{} {
	return nil
}

func (jdb *jsonDB) Create(ex tracker.Exercise) (string, error) {
	if err := jdb.db.Write(tracker.Collection, ex.GetID(), ex); err != nil {
		return "", err
	}

	return ex.GetID(), nil
}

func (jdb *jsonDB) UpdateOne(id string, object interface{}) bool {
	return false
}
