package event

import (
    "context"

    "github.com/UCMercedACM/Tenaya/pkg/entities"
)

// Repository interface allows us to access the CRUD Operations in postgresql here.
type Repository interface {
    ReadEvents() (*[]entities.Event, error)
    ReadEventsByType(eventType string) (*[]entities.Event, error)
    ReadEventByID(ID int64) (*entities.Event, error)
    CreateEvent(event *entities.Event) (*entities.Event, error)
    UpdateEvent(event  *entities.Event) (*entities.Event, error)
    DeleteEvent(ID int64) error
}

func repository  struct  {
    
}
