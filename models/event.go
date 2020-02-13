package models

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
func (event *Event) GetEvents() map[string]interface{} {
	return event
}

// GetEvents --> Returns list/map of all events based off the type of events requested
func (event *Event) GetEvents(eventType string) map[string]interface{} {
	return event
}

// GetEvent --> Returns specific event based off the id number
func (event *Event) GetEvent(id int)map[string]interface{} {
	return event
}
