package sqlite

import (
	"database/sql"
	"deni1688/exerciseTracker/config"
	"deni1688/exerciseTracker/exercise"
	"errors"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type exerciseRepository struct {
	db *sql.DB
}

func (r *exerciseRepository) FindAll() (*[]exercise.Exercise, error) {
	results := make([]exercise.Exercise, 0)

	rows, err := r.db.Query("SELECT * FROM exercises")
	if err != nil {
		return nil, errors.New("error selecting exercises: " + err.Error())
	}

	for rows.Next() {
		ex := exercise.Exercise{}
		_ = rows.Scan(
			&ex.ID,
			&ex.Name,
			&ex.Category,
			&ex.Weight,
			&ex.Duration,
			&ex.Distance,
			&ex.Reps,
			&ex.Sets,
			&ex.Created,
		)

		results = append(results, ex)
	}

	return &results, err
}

func (r *exerciseRepository) Create(ex *exercise.Exercise) (string, error) {
	ex.ID = uuid.New().String()
	ex.Created = time.Now().Unix()

	stmt, err := r.db.Prepare(`
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

func NewExerciseRepository() (exercise.Repository, error) {
	dir := config.GetString("storage.path")

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
