package jsondb

import (
	"deni1688/exerciseTracker/exercises"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
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

func (jdb *exerciseRepository) FindAll() (*[]exercises.Exercise, error) {
	return nil, nil
}

func (jdb *exerciseRepository) FindOne(id string) (*exercises.Exercise, error) {
	return nil, nil
}

func (jdb *exerciseRepository) Create(ex *exercises.Exercise) (string, error) {
	ex.ID = uuid.New().String()
	ex.Created = time.Now().Unix()

	if err := jdb.db.Write(exercises.Collection, ex.ID, ex); err != nil {
		return "", err
	}

	return ex.ID, nil
}

func (jdb *exerciseRepository) UpdateOne(id string, ex *exercises.Exercise) (bool, error) {
	return false, nil
}
