package sqlite

import (
	"database/sql"
	"deni1688/exerciseTracker/exercises"
	"errors"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

type exerciseRepository struct {
	db *sql.DB
}

func (e *exerciseRepository) FindAll() (*[]exercises.Exercise, error) {
	panic("implement me")
}

func (e *exerciseRepository) FindOne(s string) (*exercises.Exercise, error) {
	panic("implement me")
}

func (e *exerciseRepository) Create(ex *exercises.Exercise) (string, error) {
	ex.ID = uuid.New().String()
	ex.Created = time.Now().Unix()

	stmt, err := e.db.Prepare(`
		INSERT INTO exercises 
		(id, category, name, weight, duration, distance, reps, sets, created) 
		VALUES (?,?,?,?,?,?,?,?, ?)
    `)
	_, err = stmt.Exec(
		ex.ID,
		ex.Category,
		ex.Name,
		ex.Weight,
		ex.Duration,
		ex.Distance,
		ex.Reps,
		ex.Sets,
		ex.Created,
	)
	if err != nil {
		return "", errors.New("error creating exercise: " + err.Error())
	}

	return ex.ID, nil
}

func (e *exerciseRepository) UpdateOne(s string, exercise *exercises.Exercise) (bool, error) {
	panic("implement me")
}

func NewExerciseRepository() (exercises.Repository, error) {
	dir := os.Getenv("SQLITE_DB_DIR")

	db, err := sql.Open("sqlite3", dir)
	if err != nil {
		return nil, errors.New("error connecting sql database: " + err.Error())
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS exercises (
			id VARCHAR(32) PRIMARY KEY,
		 	name TEXT,
		    category TEXT,
		    weight FLOAT,
		    duration INT,
        	distance INT,
            reps INT,
            sets INT,
            created INT)`)

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("error creating exercises table: " + err.Error())
	}
	return &exerciseRepository{db}, nil
}
