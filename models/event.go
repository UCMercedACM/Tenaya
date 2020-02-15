package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Time        string `json:"time"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

// GetEvents --> Returns list/map of all events
func (event *Event) GetAllEvents() {
	return event
}

// GetEvents --> Returns list/map of all events based off the type of events requested
func (event *Event) GetEvents(eventType string) Event {
	return event
}

// GetEvent --> Returns specific event based off the id number
func (event *Event) GetEvent(id int) map[string]interface{} {
	return event
}

// CreateEvent --> Creates a new event
func (event *Event) CreateEvent() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
