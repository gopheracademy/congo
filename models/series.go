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

// This is the Series model
type Series struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	CreatedAt time.Time
	DeletedAt *time.Time
	Name      string
	TenantID  int // Belongs To Tenant
	UpdatedAt time.Time
	Tenant    Tenant
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Series) TableName() string {
	return "series"

}

// SeriesDB is the implementation of the storage interface for
// Series.
type SeriesDB struct {
	Db gorm.DB
}

// NewSeriesDB creates a new storage type.
func NewSeriesDB(db gorm.DB) *SeriesDB {
	return &SeriesDB{Db: db}
}

// DB returns the underlying database.
func (m *SeriesDB) DB() interface{} {
	return &m.Db
}

// SeriesStorage represents the storage interface.
type SeriesStorage interface {
	DB() interface{}
	List(ctx context.Context) []Series
	Get(ctx context.Context, id int) (Series, error)
	Add(ctx context.Context, series *Series) (*Series, error)
	Update(ctx context.Context, series *Series) error
	Delete(ctx context.Context, id int) error

	ListSeries(ctx context.Context, tenantID int) []*app.Series
	OneSeries(ctx context.Context, id int, tenantID int) (*app.Series, error)

	UpdateFromCreateSeriesPayload(ctx context.Context, payload *app.CreateSeriesPayload, id int) error

	UpdateFromUpdateSeriesPayload(ctx context.Context, payload *app.UpdateSeriesPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *SeriesDB) TableName() string {
	return "series"

}

// Belongs To Relationships

// SeriesFilterByTenant is a gorm filter for a Belongs To relationship.
func SeriesFilterByTenant(tenantID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if tenantID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("tenant_id = ?", tenantID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Series as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *SeriesDB) Get(ctx context.Context, id int) (Series, error) {
	defer goa.MeasureSince([]string{"goa", "db", "series", "get"}, time.Now())

	var native Series
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Series{}, nil
	}

	return native, err
}

// List returns an array of Series
func (m *SeriesDB) List(ctx context.Context) []Series {
	defer goa.MeasureSince([]string{"goa", "db", "series", "list"}, time.Now())

	var objs []Series
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Series", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *SeriesDB) Add(ctx context.Context, model *Series) (*Series, error) {
	defer goa.MeasureSince([]string{"goa", "db", "series", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Series", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *SeriesDB) Update(ctx context.Context, model *Series) error {
	defer goa.MeasureSince([]string{"goa", "db", "series", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *SeriesDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "series", "delete"}, time.Now())

	var obj Series

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Series", "error", err.Error())
		return err
	}

	return nil
}

// SeriesFromCreateSeriesPayload Converts source CreateSeriesPayload to target Series model
// only copying the non-nil fields from the source.
func SeriesFromCreateSeriesPayload(payload *app.CreateSeriesPayload) *Series {
	series := &Series{}
	series.Name = payload.Name

	return series
}

// UpdateFromCreateSeriesPayload applies non-nil changes from CreateSeriesPayload to the model and saves it
func (m *SeriesDB) UpdateFromCreateSeriesPayload(ctx context.Context, payload *app.CreateSeriesPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "series", "updatefromcreateSeriesPayload"}, time.Now())

	var obj Series
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Series", "error", err.Error())
		return err
	}
	obj.Name = payload.Name

	err = m.Db.Save(&obj).Error
	return err
}

// SeriesFromUpdateSeriesPayload Converts source UpdateSeriesPayload to target Series model
// only copying the non-nil fields from the source.
func SeriesFromUpdateSeriesPayload(payload *app.UpdateSeriesPayload) *Series {
	series := &Series{}
	if payload.Name != nil {
		series.Name = *payload.Name
	}

	return series
}

// UpdateFromUpdateSeriesPayload applies non-nil changes from UpdateSeriesPayload to the model and saves it
func (m *SeriesDB) UpdateFromUpdateSeriesPayload(ctx context.Context, payload *app.UpdateSeriesPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "series", "updatefromupdateSeriesPayload"}, time.Now())

	var obj Series
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Series", "error", err.Error())
		return err
	}
	if payload.Name != nil {
		obj.Name = *payload.Name
	}

	err = m.Db.Save(&obj).Error
	return err
}
