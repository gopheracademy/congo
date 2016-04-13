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

// This is the Tenant model
type Tenant struct {
	ID        int `gorm:"primary_key"` // This is the ID PK field
	CreatedAt time.Time
	DeletedAt *time.Time
	Events    []Event // has many Events
	Name      string
	UpdatedAt time.Time
	Users     []User // has many Users
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Tenant) TableName() string {
	return "tenants"

}

// TenantDB is the implementation of the storage interface for
// Tenant.
type TenantDB struct {
	Db gorm.DB
}

// NewTenantDB creates a new storage type.
func NewTenantDB(db gorm.DB) *TenantDB {
	return &TenantDB{Db: db}
}

// DB returns the underlying database.
func (m *TenantDB) DB() interface{} {
	return &m.Db
}

// TenantStorage represents the storage interface.
type TenantStorage interface {
	DB() interface{}
	List(ctx context.Context) []Tenant
	Get(ctx context.Context, id int) (Tenant, error)
	Add(ctx context.Context, tenant *Tenant) (*Tenant, error)
	Update(ctx context.Context, tenant *Tenant) error
	Delete(ctx context.Context, id int) error

	ListTenant(ctx context.Context) []*app.Tenant
	OneTenant(ctx context.Context, id int) (*app.Tenant, error)

	UpdateFromCreateTenantPayload(ctx context.Context, payload *app.CreateTenantPayload, id int) error

	UpdateFromUpdateTenantPayload(ctx context.Context, payload *app.UpdateTenantPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *TenantDB) TableName() string {
	return "tenants"

}

// CRUD Functions

// Get returns a single Tenant as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *TenantDB) Get(ctx context.Context, id int) (Tenant, error) {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "get"}, time.Now())

	var native Tenant
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Tenant{}, nil
	}

	return native, err
}

// List returns an array of Tenant
func (m *TenantDB) List(ctx context.Context) []Tenant {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "list"}, time.Now())

	var objs []Tenant
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Tenant", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *TenantDB) Add(ctx context.Context, model *Tenant) (*Tenant, error) {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Tenant", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *TenantDB) Update(ctx context.Context, model *Tenant) error {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *TenantDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "delete"}, time.Now())

	var obj Tenant

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Tenant", "error", err.Error())
		return err
	}

	return nil
}

// TenantFromCreateTenantPayload Converts source CreateTenantPayload to target Tenant model
// only copying the non-nil fields from the source.
func TenantFromCreateTenantPayload(payload *app.CreateTenantPayload) *Tenant {
	tenant := &Tenant{}
	tenant.Name = payload.Name

	return tenant
}

// UpdateFromCreateTenantPayload applies non-nil changes from CreateTenantPayload to the model and saves it
func (m *TenantDB) UpdateFromCreateTenantPayload(ctx context.Context, payload *app.CreateTenantPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "updatefromcreateTenantPayload"}, time.Now())

	var obj Tenant
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Tenant", "error", err.Error())
		return err
	}
	obj.Name = payload.Name

	err = m.Db.Save(&obj).Error
	return err
}

// TenantFromUpdateTenantPayload Converts source UpdateTenantPayload to target Tenant model
// only copying the non-nil fields from the source.
func TenantFromUpdateTenantPayload(payload *app.UpdateTenantPayload) *Tenant {
	tenant := &Tenant{}
	if payload.Name != nil {
		tenant.Name = *payload.Name
	}

	return tenant
}

// UpdateFromUpdateTenantPayload applies non-nil changes from UpdateTenantPayload to the model and saves it
func (m *TenantDB) UpdateFromUpdateTenantPayload(ctx context.Context, payload *app.UpdateTenantPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "updatefromupdateTenantPayload"}, time.Now())

	var obj Tenant
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Tenant", "error", err.Error())
		return err
	}
	if payload.Name != nil {
		obj.Name = *payload.Name
	}

	err = m.Db.Save(&obj).Error
	return err
}
