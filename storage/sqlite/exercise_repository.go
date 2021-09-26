package sqlite

import (
	"database/sql"
	"deni1688/exercise_tracker/config"
	"deni1688/exercise_tracker/domain"
	"errors"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type exerciseRepository struct {
	db *sql.DB
}

func NewExerciseRepository() (domain.ExerciseRepository, error) {
	dir := config.GetString("storage.path")

	db, err := sql.Open("sqlite3", dir)
	if err != nil {
		return nil, errors.New("error connecting sql database: " + err.Error())
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS exercises (
			id VARCHAR(32) PRIMARY KEY,
			schedule_id VARCHAR(32),
		 	name TEXT,
		    category TEXT,
		    weight FLOAT,
		    minute_duration FLOAT,
        	km_distance INT,
            reps INT,
            sets INT,
            created DATE DEFAULT CURRENT_TIMESTAMP)`)

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("error creating exercises table: " + err.Error())
	}
	return &exerciseRepository{db}, nil
}

func (r *exerciseRepository) FindAll() (*[]domain.Exercise, error) {
	results := make([]domain.Exercise, 0)

	rows, err := r.db.Query("SELECT * FROM exercises")
	if err != nil {
		return nil, errors.New("error selecting exercises: " + err.Error())
	}

	for rows.Next() {
		ex := domain.Exercise{}
		_ = rows.Scan(
			&ex.ID,
			&ex.Name,
			&ex.Category,
			&ex.Weight,
			&ex.MinuteDuration,
			&ex.KmDistance,
			&ex.Reps,
			&ex.Sets,
			&ex.Created,
		)

		results = append(results, ex)
	}

	return &results, err
}

func (r *exerciseRepository) Create(ex *domain.Exercise) (string, error) {
	ex.ID = uuid.New().String()

	stmt, err := r.db.Prepare(`
		INSERT INTO exercises 
		(id, schedule_id, category, name, weight, minute_duration, km_distance, reps, sets) 
		VALUES (?,?,?,?,?,?,?,?,?)
    `)
	if err != nil {
		return "", errors.New("error creating exercise: " + err.Error())
	}
	_, err = stmt.Exec(
		ex.ID,
		ex.ScheduleID,
		ex.Category,
		ex.Name,
		ex.Weight,
		ex.MinuteDuration,
		ex.KmDistance,
		ex.Reps,
		ex.Sets,
	)
	if err != nil {
		return "", errors.New("error creating exercise: " + err.Error())
	}

	return ex.ID, nil
}
