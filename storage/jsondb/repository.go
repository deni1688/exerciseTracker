package jsondb

import (
	"deni1688/myHealthTrack/tracker"
	"encoding/json"
	"fmt"
	"log"

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

func (jdb *jsonDB) FindAll() ([]tracker.TrackEntry, error) {
	records, err := jdb.db.ReadAll("track_entries")
	if err != nil {
		return nil, err
	}

	trackEntries := make([]tracker.TrackEntry, 0)
	for _, record := range records {
		trackEntryFound := tracker.TrackEntry{}
		if err := json.Unmarshal([]byte(record), &trackEntryFound); err != nil {
			log.Println("Error in FindAll decoding TrackEntry", err)
		}
		trackEntries = append(trackEntries, trackEntryFound)
	}

	return trackEntries, nil
}

func (jdb *jsonDB) FindOne(id string) interface{} {
	return nil
}

func (jdb *jsonDB) Create(te *tracker.TrackEntry) (string, error) {
	if err := jdb.db.Write("track_entries", te.ID, te); err != nil {
		return "", err
	}

	return te.ID, nil
}

func (jdb *jsonDB) UpdateOne(id string, object interface{}) bool {
	return false
}
