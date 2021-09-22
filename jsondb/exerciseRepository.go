package jsondb

import (
	"deni1688/exerciseTracker/exercises"
	"errors"
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
	"os"
)

type exerciseRepository struct {
	db *scribble.Driver
}

func NewExerciseRepository() (exercises.Repository, error) {
	dir := os.Getenv("JSON_DB_DIR")
	if dir == "" {
		return nil, errors.New("error getting the JSON_DB_DIR env variable")
	}

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return &exerciseRepository{db}, nil
}

func (jdb *exerciseRepository) FindAll() ([]exercises.Entity, error) {
	return nil, nil
}

func (jdb *exerciseRepository) FindOne(id string) interface{} {
	return nil
}

func (jdb *exerciseRepository) Create(ex exercises.Entity) (string, error) {
	if err := jdb.db.Write(exercises.Collection, ex.GetID(), ex); err != nil {
		return "", err
	}

	return ex.GetID(), nil
}

func (jdb *exerciseRepository) UpdateOne(id string, object interface{}) bool {
	return false
}
