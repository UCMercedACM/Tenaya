package entities

import (
	"time"
)

// Event --> Struct defining the global data structure for all events
type Event struct {
	ID          int64    `json:"id" pg:"column_name:id,pk,unique,type:uuid,default:gen_random_uuid(),notnull"`
	Name        string   `json:"name" pg:"column_name:name,type:text,notnull"`
	Description string   `json:"description" pg:"column_name:description,type:text,notnull"`
	Location    string   `json:"location" pg:"column_name:location,type:text"`
	Type        string   `json:"type" pg:"column_name:type,type:text"`
	Date        string   `json:"date" pg:"column_name:date,type:text,notnull"`
	StartTime   string   `json:"startTime" pg:"column_name:start_time,type:text"`
	EndTime     string   `json:"endTime" pg:"column_name:end_time,type:text"`
	Attendees   []string `json:"attendees" pg:"column_name:attendees,type:jsonb,array"`
	HostedBy    string   `json:"hostedBy" pg:"column_name:hosted_by,type:text,default:'Association for Computing Machinery, UC Merced',notnull"`
	Tags        []string `json:"tags" pg:"column_name:tags,type:jsonb,array"`
	Series      string   `json:"series" pg:"column_name:series,type:text"`
	Image       []byte   `json:"image" pg:"column_name:image,type:bytea"`
	Flyer       []byte   `json:"flyer" pg:"column_name:flyer,type:bytea"`
	Active      bool     `json:"active" pg:"column_name:active,type:boolean,default:false,notnull"`
}

// Events --> Struct defining  multiple events
type Events struct {
	Event []Event `json:"events"`
}

// EventModel --> Struct for defining the database schema
type EventModel struct {
	tableName struct{} `pg:"events,alias:e"`

	Event

	CreatedAt time.Time `json:"createdAt" pg:"column_name:created_at,default:now(),notnull"`
	UpdatedAt time.Time `json:"updatedAt" pg:"column_name:updated_at,default:now(),notnull"`
}

// DeleteRequest Only one struct per file should exists unless another struct is closely related with the one defined in this file.
type DeleteRequest struct {
	ID string `json:"id" pg:"column_name:id,pk,unique,type:uuid,default:gen_random_uuid(),notnull"`
}
