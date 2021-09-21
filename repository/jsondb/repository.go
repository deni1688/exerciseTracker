package jsondb

import (
	"deni1688/exerciseTracker/exercises"
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
)

type jsonDB struct {
	db *scribble.Driver
}

func New(dir string) exercises.Repository {
	db, err := scribble.New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	return &jsonDB{db}
}

func (jdb *jsonDB) FindAll() ([]exercises.Entity, error) {
	return nil, nil
}

func (jdb *jsonDB) FindOne(id string) interface{} {
	return nil
}

func (jdb *jsonDB) Create(ex exercises.Entity) (string, error) {
	if err := jdb.db.Write(exercises.Collection, ex.GetID(), ex); err != nil {
		return "", err
	}

	return ex.GetID(), nil
}

func (jdb *jsonDB) UpdateOne(id string, object interface{}) bool {
	return false
}
