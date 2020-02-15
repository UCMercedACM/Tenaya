package models

import (
	"github.com/go-pg/pg/v9"
)

type Event struct {
	ID int64
	Key int64
	Name string
	Description string
	Type string
	Date string
	StartTime string
	EndTime string
}

// GetEvents --> Returns list/map of all events
func GetAllEvents() ([]Event, error) {
	var events []Event
	_, err := db.Query(&events, `SELECT * FROM events`)
	return events, err
}

// GetEvents --> Returns list/map of all events based off the type of events requested
func GetEvents(eventType string) ([]Event, error) {
	var events Event
	_, err := db.QueryOne(&events, `SELECT * FROM events WHERE id = ?`, id)
	return events, err
}

// GetEvent --> Returns specific event based off the id number
func GetEvent(db *pg.DB, id int64) (*Event, error) {
	var event Event
	_, err := db.QueryOne(&event, `SELECT * FROM events WHERE id = ?`, id)
	return &event, err
}

func CreateEvent(db *pg.DB, event *Event) error {
	_, err := db.QueryOne(event, `
		INSERT INTO users (name, emails) VALUES (?name, ?emails)
		RETURNING id
	`, event)
	return err
}
