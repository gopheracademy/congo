//************************************************************************//
// API "congo": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// This is the Event model
type Event struct {
	ID            int `gorm:"primary_key"` // This is the ID PK field
	CreatedAt     time.Time
	DeletedAt     *time.Time
	EndDate       *time.Time
	Name          string
	Presentations []Presentation // has many Presentations
	Speakers      []Speaker      // has many Speakers
	StartDate     *time.Time
	TenantID      int // has many Event
	URL           *string
	UpdatedAt     time.Time
	Tenant        Tenant
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Event) TableName() string {
	return "events"

}

// EventDB is the implementation of the storage interface for
// Event.
type EventDB struct {
	Db gorm.DB
}

// NewEventDB creates a new storage type.
func NewEventDB(db gorm.DB) *EventDB {
	return &EventDB{Db: db}
}

// DB returns the underlying database.
func (m *EventDB) DB() interface{} {
	return &m.Db
}

// EventStorage represents the storage interface.
type EventStorage interface {
	DB() interface{}
	List(ctx context.Context) []Event
	Get(ctx context.Context, id int) (Event, error)
	Add(ctx context.Context, event *Event) (*Event, error)
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id int) error

	ListEvent(ctx context.Context, tenantID int) []*app.Event
	OneEvent(ctx context.Context, id int, tenantID int) (*app.Event, error)

	UpdateFromCreateEventPayload(ctx context.Context, payload *app.CreateEventPayload, id int) error

	UpdateFromUpdateEventPayload(ctx context.Context, payload *app.UpdateEventPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *EventDB) TableName() string {
	return "events"

}

// Belongs To Relationships

// EventFilterByTenant is a gorm filter for a Belongs To relationship.
func EventFilterByTenant(tenantID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if tenantID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("tenant_id = ?", tenantID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Event as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *EventDB) Get(ctx context.Context, id int) (Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "get"}, time.Now())

	var native Event
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Event{}, nil
	}

	return native, err
}

// List returns an array of Event
func (m *EventDB) List(ctx context.Context) []Event {
	defer goa.MeasureSince([]string{"goa", "db", "event", "list"}, time.Now())

	var objs []Event
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Event", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *EventDB) Add(ctx context.Context, model *Event) (*Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Event", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *EventDB) Update(ctx context.Context, model *Event) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *EventDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "delete"}, time.Now())

	var obj Event

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Event", "error", err.Error())
		return err
	}

	return nil
}

// EventFromCreateEventPayload Converts source CreateEventPayload to target Event model
// only copying the non-nil fields from the source.
func EventFromCreateEventPayload(payload *app.CreateEventPayload) *Event {
	event := &Event{}
	if payload.EndDate != nil {
		event.EndDate = payload.EndDate
	}
	event.Name = payload.Name
	if payload.StartDate != nil {
		event.StartDate = payload.StartDate
	}
	if payload.URL != nil {
		event.URL = payload.URL
	}

	return event
}

// UpdateFromCreateEventPayload applies non-nil changes from CreateEventPayload to the model and saves it
func (m *EventDB) UpdateFromCreateEventPayload(ctx context.Context, payload *app.CreateEventPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "updatefromcreateEventPayload"}, time.Now())

	var obj Event
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Event", "error", err.Error())
		return err
	}
	if payload.EndDate != nil {
		obj.EndDate = payload.EndDate
	}
	obj.Name = payload.Name
	if payload.StartDate != nil {
		obj.StartDate = payload.StartDate
	}
	if payload.URL != nil {
		obj.URL = payload.URL
	}

	err = m.Db.Save(&obj).Error
	return err
}

// EventFromUpdateEventPayload Converts source UpdateEventPayload to target Event model
// only copying the non-nil fields from the source.
func EventFromUpdateEventPayload(payload *app.UpdateEventPayload) *Event {
	event := &Event{}
	if payload.EndDate != nil {
		event.EndDate = payload.EndDate
	}
	if payload.Name != nil {
		event.Name = *payload.Name
	}
	if payload.StartDate != nil {
		event.StartDate = payload.StartDate
	}
	if payload.URL != nil {
		event.URL = payload.URL
	}

	return event
}

// UpdateFromUpdateEventPayload applies non-nil changes from UpdateEventPayload to the model and saves it
func (m *EventDB) UpdateFromUpdateEventPayload(ctx context.Context, payload *app.UpdateEventPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "event", "updatefromupdateEventPayload"}, time.Now())

	var obj Event
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Event", "error", err.Error())
		return err
	}
	if payload.EndDate != nil {
		obj.EndDate = payload.EndDate
	}
	if payload.Name != nil {
		obj.Name = *payload.Name
	}
	if payload.StartDate != nil {
		obj.StartDate = payload.StartDate
	}
	if payload.URL != nil {
		obj.URL = payload.URL
	}

	err = m.Db.Save(&obj).Error
	return err
}
