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

// This is the presentation model
type Presentation struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	Abstract  string
	CreatedAt time.Time
	DeletedAt *time.Time
	Detail    *string
	EventID   int // Belongs To Event
	Name      *string
	SpeakerID int // Belongs To Speaker
	UpdatedAt time.Time
	Event     Event
	Speaker   Speaker
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Presentation) TableName() string {
	return "presentations"

}

// PresentationDB is the implementation of the storage interface for
// Presentation.
type PresentationDB struct {
	Db gorm.DB
}

// NewPresentationDB creates a new storage type.
func NewPresentationDB(db gorm.DB) *PresentationDB {
	return &PresentationDB{Db: db}
}

// DB returns the underlying database.
func (m *PresentationDB) DB() interface{} {
	return &m.Db
}

// PresentationStorage represents the storage interface.
type PresentationStorage interface {
	DB() interface{}
	List(ctx context.Context) []Presentation
	Get(ctx context.Context, id int) (Presentation, error)
	Add(ctx context.Context, presentation *Presentation) (*Presentation, error)
	Update(ctx context.Context, presentation *Presentation) error
	Delete(ctx context.Context, id int) error

	ListPresentationAdmin(ctx context.Context, eventID int, speakerID int) []*app.PresentationAdmin
	OnePresentationAdmin(ctx context.Context, id int, eventID int, speakerID int) (*app.PresentationAdmin, error)

	ListPresentation(ctx context.Context, eventID int, speakerID int) []*app.Presentation
	OnePresentation(ctx context.Context, id int, eventID int, speakerID int) (*app.Presentation, error)

	UpdateFromCreatePresentationPayload(ctx context.Context, payload *app.CreatePresentationPayload, id int) error

	UpdateFromUpdatePresentationPayload(ctx context.Context, payload *app.UpdatePresentationPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PresentationDB) TableName() string {
	return "presentations"

}

// Belongs To Relationships

// PresentationFilterByEvent is a gorm filter for a Belongs To relationship.
func PresentationFilterByEvent(eventID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if eventID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("event_id = ?", eventID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// PresentationFilterBySpeaker is a gorm filter for a Belongs To relationship.
func PresentationFilterBySpeaker(speakerID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if speakerID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("speaker_id = ?", speakerID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Presentation as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PresentationDB) Get(ctx context.Context, id int) (Presentation, error) {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "get"}, time.Now())

	var native Presentation
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Presentation{}, nil
	}

	return native, err
}

// List returns an array of Presentation
func (m *PresentationDB) List(ctx context.Context) []Presentation {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "list"}, time.Now())

	var objs []Presentation
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Presentation", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *PresentationDB) Add(ctx context.Context, model *Presentation) (*Presentation, error) {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Presentation", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *PresentationDB) Update(ctx context.Context, model *Presentation) error {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PresentationDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "delete"}, time.Now())

	var obj Presentation

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Presentation", "error", err.Error())
		return err
	}

	return nil
}

// PresentationFromCreatePresentationPayload Converts source CreatePresentationPayload to target Presentation model
// only copying the non-nil fields from the source.
func PresentationFromCreatePresentationPayload(payload *app.CreatePresentationPayload) *Presentation {
	presentation := &Presentation{}
	presentation.Abstract = payload.Abstract
	if payload.Detail != nil {
		presentation.Detail = payload.Detail
	}
	if payload.Name != nil {
		presentation.Name = payload.Name
	}

	return presentation
}

// UpdateFromCreatePresentationPayload applies non-nil changes from CreatePresentationPayload to the model and saves it
func (m *PresentationDB) UpdateFromCreatePresentationPayload(ctx context.Context, payload *app.CreatePresentationPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "updatefromcreatePresentationPayload"}, time.Now())

	var obj Presentation
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Presentation", "error", err.Error())
		return err
	}
	obj.Abstract = payload.Abstract
	if payload.Detail != nil {
		obj.Detail = payload.Detail
	}
	if payload.Name != nil {
		obj.Name = payload.Name
	}

	err = m.Db.Save(&obj).Error
	return err
}

// PresentationFromUpdatePresentationPayload Converts source UpdatePresentationPayload to target Presentation model
// only copying the non-nil fields from the source.
func PresentationFromUpdatePresentationPayload(payload *app.UpdatePresentationPayload) *Presentation {
	presentation := &Presentation{}
	if payload.Abstract != nil {
		presentation.Abstract = *payload.Abstract
	}
	if payload.Detail != nil {
		presentation.Detail = payload.Detail
	}
	if payload.Name != nil {
		presentation.Name = payload.Name
	}

	return presentation
}

// UpdateFromUpdatePresentationPayload applies non-nil changes from UpdatePresentationPayload to the model and saves it
func (m *PresentationDB) UpdateFromUpdatePresentationPayload(ctx context.Context, payload *app.UpdatePresentationPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "updatefromupdatePresentationPayload"}, time.Now())

	var obj Presentation
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Presentation", "error", err.Error())
		return err
	}
	if payload.Abstract != nil {
		obj.Abstract = *payload.Abstract
	}
	if payload.Detail != nil {
		obj.Detail = payload.Detail
	}
	if payload.Name != nil {
		obj.Name = payload.Name
	}

	err = m.Db.Save(&obj).Error
	return err
}
