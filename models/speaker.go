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

// This is the Speaker model
type Speaker struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	Bio       *string
	CreatedAt time.Time
	DeletedAt *time.Time
	EventID   int // Belongs To Event
	FirstName string
	Github    *string
	ImageURL  *string
	LastName  string
	Linkedin  *string
	Twitter   *string
	UpdatedAt time.Time
	Event     Event
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Speaker) TableName() string {
	return "speakers"

}

// SpeakerDB is the implementation of the storage interface for
// Speaker.
type SpeakerDB struct {
	Db gorm.DB
}

// NewSpeakerDB creates a new storage type.
func NewSpeakerDB(db gorm.DB) *SpeakerDB {
	return &SpeakerDB{Db: db}
}

// DB returns the underlying database.
func (m *SpeakerDB) DB() interface{} {
	return &m.Db
}

// SpeakerStorage represents the storage interface.
type SpeakerStorage interface {
	DB() interface{}
	List(ctx context.Context) []Speaker
	Get(ctx context.Context, id int) (Speaker, error)
	Add(ctx context.Context, speaker *Speaker) (*Speaker, error)
	Update(ctx context.Context, speaker *Speaker) error
	Delete(ctx context.Context, id int) error

	ListSpeakerAdmin(ctx context.Context, eventID int) []*app.SpeakerAdmin
	OneSpeakerAdmin(ctx context.Context, id int, eventID int) (*app.SpeakerAdmin, error)

	ListSpeaker(ctx context.Context, eventID int) []*app.Speaker
	OneSpeaker(ctx context.Context, id int, eventID int) (*app.Speaker, error)

	ListSpeakerLink(ctx context.Context, eventID int) []*app.SpeakerLink
	OneSpeakerLink(ctx context.Context, id int, eventID int) (*app.SpeakerLink, error)

	UpdateFromCreateSpeakerPayload(ctx context.Context, payload *app.CreateSpeakerPayload, id int) error

	UpdateFromUpdateSpeakerPayload(ctx context.Context, payload *app.UpdateSpeakerPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *SpeakerDB) TableName() string {
	return "speakers"

}

// Belongs To Relationships

// SpeakerFilterByEvent is a gorm filter for a Belongs To relationship.
func SpeakerFilterByEvent(eventID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if eventID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("event_id = ?", eventID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Speaker as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *SpeakerDB) Get(ctx context.Context, id int) (Speaker, error) {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "get"}, time.Now())

	var native Speaker
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Speaker{}, nil
	}

	return native, err
}

// List returns an array of Speaker
func (m *SpeakerDB) List(ctx context.Context) []Speaker {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "list"}, time.Now())

	var objs []Speaker
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Speaker", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *SpeakerDB) Add(ctx context.Context, model *Speaker) (*Speaker, error) {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Speaker", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *SpeakerDB) Update(ctx context.Context, model *Speaker) error {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *SpeakerDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "delete"}, time.Now())

	var obj Speaker

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Speaker", "error", err.Error())
		return err
	}

	return nil
}

// SpeakerFromCreateSpeakerPayload Converts source CreateSpeakerPayload to target Speaker model
// only copying the non-nil fields from the source.
func SpeakerFromCreateSpeakerPayload(payload *app.CreateSpeakerPayload) *Speaker {
	speaker := &Speaker{}
	if payload.Bio != nil {
		speaker.Bio = payload.Bio
	}
	speaker.FirstName = payload.FirstName
	if payload.Github != nil {
		speaker.Github = payload.Github
	}
	if payload.ImageURL != nil {
		speaker.ImageURL = payload.ImageURL
	}
	speaker.LastName = payload.LastName
	if payload.Linkedin != nil {
		speaker.Linkedin = payload.Linkedin
	}
	if payload.Twitter != nil {
		speaker.Twitter = payload.Twitter
	}

	return speaker
}

// UpdateFromCreateSpeakerPayload applies non-nil changes from CreateSpeakerPayload to the model and saves it
func (m *SpeakerDB) UpdateFromCreateSpeakerPayload(ctx context.Context, payload *app.CreateSpeakerPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "updatefromcreateSpeakerPayload"}, time.Now())

	var obj Speaker
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Speaker", "error", err.Error())
		return err
	}
	if payload.Bio != nil {
		obj.Bio = payload.Bio
	}
	obj.FirstName = payload.FirstName
	if payload.Github != nil {
		obj.Github = payload.Github
	}
	if payload.ImageURL != nil {
		obj.ImageURL = payload.ImageURL
	}
	obj.LastName = payload.LastName
	if payload.Linkedin != nil {
		obj.Linkedin = payload.Linkedin
	}
	if payload.Twitter != nil {
		obj.Twitter = payload.Twitter
	}

	err = m.Db.Save(&obj).Error
	return err
}

// SpeakerFromUpdateSpeakerPayload Converts source UpdateSpeakerPayload to target Speaker model
// only copying the non-nil fields from the source.
func SpeakerFromUpdateSpeakerPayload(payload *app.UpdateSpeakerPayload) *Speaker {
	speaker := &Speaker{}
	if payload.Bio != nil {
		speaker.Bio = payload.Bio
	}
	if payload.FirstName != nil {
		speaker.FirstName = *payload.FirstName
	}
	if payload.Github != nil {
		speaker.Github = payload.Github
	}
	if payload.ImageURL != nil {
		speaker.ImageURL = payload.ImageURL
	}
	if payload.LastName != nil {
		speaker.LastName = *payload.LastName
	}
	if payload.Linkedin != nil {
		speaker.Linkedin = payload.Linkedin
	}
	if payload.Twitter != nil {
		speaker.Twitter = payload.Twitter
	}

	return speaker
}

// UpdateFromUpdateSpeakerPayload applies non-nil changes from UpdateSpeakerPayload to the model and saves it
func (m *SpeakerDB) UpdateFromUpdateSpeakerPayload(ctx context.Context, payload *app.UpdateSpeakerPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "updatefromupdateSpeakerPayload"}, time.Now())

	var obj Speaker
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Speaker", "error", err.Error())
		return err
	}
	if payload.Bio != nil {
		obj.Bio = payload.Bio
	}
	if payload.FirstName != nil {
		obj.FirstName = *payload.FirstName
	}
	if payload.Github != nil {
		obj.Github = payload.Github
	}
	if payload.ImageURL != nil {
		obj.ImageURL = payload.ImageURL
	}
	if payload.LastName != nil {
		obj.LastName = *payload.LastName
	}
	if payload.Linkedin != nil {
		obj.Linkedin = payload.Linkedin
	}
	if payload.Twitter != nil {
		obj.Twitter = payload.Twitter
	}

	err = m.Db.Save(&obj).Error
	return err
}
