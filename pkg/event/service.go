package event

import "github.com/UCMercedACM/Tenaya/pkg/entities"

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
    FetchEvents() (*[]entities.Event, error)
    FetchEventsByType(eventType string)  (*[]entities.Event, error)
    FetchEventByID(ID int64) (*entities.Event, error)
	InsertEvent(event *entities.Event) (*entities.Event, error)
	UpdateEvent(event *entities.Event) (*entities.Event, error)
	RemoveEvent(ID int64) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchEvents() (*[]entities.Event, error) {
	return s.repository.ReadEvents()
}

func (s *service) FetchEventsByType(eventType string) (*[]entities.Event, error) {
	return s.repository.ReadEventsByType(eventType)
}

func (s *service) FetchEventByID(ID int64) (*entities.Event, error) {
	return s.repository.ReadEventByID(ID)
}

func (s *service) InsertEvent(event *entities.Event) (*entities.Event, error) {
	return s.repository.CreateEvent(event)
}

func (s *service) UpdateEvent(event *entities.Event) (*entities.Event, error) {
	return s.repository.UpdateEvent(event)
}

func (s *service) RemoveEvent(ID int64) error {
	return s.repository.DeleteEvent(ID)
}
