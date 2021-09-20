package tracker

type Repository interface {
	FindAll() ([]TrackEntry, error)
	FindOne(id string) interface{}
	Create(te *TrackEntry) (string, error)
	UpdateOne(id string, object interface{}) bool
}
