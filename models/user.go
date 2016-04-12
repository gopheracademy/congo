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

// This is the User model
type User struct {
	ID             int `gorm:"primary_key"` // This is the ID PK field
	CreatedAt      time.Time
	DeletedAt      *time.Time
	Email          string
	FirstName      string `gorm:"column:firstname"`
	Href           string
	LastName       string `gorm:"column:lastname"`
	MemberID       *int
	Password       string
	Role           string
	TenantID       int // Belongs To Tenant
	UpdatedAt      time.Time
	Validated      *bool
	ValidationCode *string
	Tenant         Tenant
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	Get(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error

	ListUser(ctx context.Context, tenantID int) []*app.User
	OneUser(ctx context.Context, id int, tenantID int) (*app.User, error)

	ListUserLink(ctx context.Context, tenantID int) []*app.UserLink
	OneUserLink(ctx context.Context, id int, tenantID int) (*app.UserLink, error)

	ListUserTenant(ctx context.Context, tenantID int) []*app.UserTenant
	OneUserTenant(ctx context.Context, id int, tenantID int) (*app.UserTenant, error)

	UpdateFromCreateAdminuserPayload(ctx context.Context, payload *app.CreateAdminuserPayload, id int) error

	UpdateFromCreateUserPayload(ctx context.Context, payload *app.CreateUserPayload, id int) error

	UpdateFromUpdateAdminuserPayload(ctx context.Context, payload *app.UpdateAdminuserPayload, id int) error

	UpdateFromUpdateUserPayload(ctx context.Context, payload *app.UpdateUserPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// Belongs To Relationships

// UserFilterByTenant is a gorm filter for a Belongs To relationship.
func UserFilterByTenant(tenantID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if tenantID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("tenant_id = ?", tenantID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id int) (User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "get"}, time.Now())

	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, nil
	}

	return native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) []User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "list"}, time.Now())

	var objs []User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *UserDB) Add(ctx context.Context, model *User) (*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating User", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "delete"}, time.Now())

	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}

	return nil
}

// UserFromCreateAdminuserPayload Converts source CreateAdminuserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromCreateAdminuserPayload(payload *app.CreateAdminuserPayload) *User {
	user := &User{}
	user.Email = payload.Email
	user.FirstName = payload.FirstName
	if payload.ID != nil {
		user.ID = *payload.ID
	}
	user.LastName = payload.LastName
	if payload.MemberID != nil {
		user.MemberID = payload.MemberID
	}
	user.Password = payload.Password
	user.Role = payload.Role
	if payload.Validated != nil {
		user.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		user.ValidationCode = payload.ValidationCode
	}

	return user
}

// UpdateFromCreateAdminuserPayload applies non-nil changes from CreateAdminuserPayload to the model and saves it
func (m *UserDB) UpdateFromCreateAdminuserPayload(ctx context.Context, payload *app.CreateAdminuserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromcreateAdminuserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	obj.Email = payload.Email
	obj.FirstName = payload.FirstName
	if payload.ID != nil {
		obj.ID = *payload.ID
	}
	obj.LastName = payload.LastName
	if payload.MemberID != nil {
		obj.MemberID = payload.MemberID
	}
	obj.Password = payload.Password
	obj.Role = payload.Role
	if payload.Validated != nil {
		obj.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		obj.ValidationCode = payload.ValidationCode
	}

	err = m.Db.Save(&obj).Error
	return err
}

// UserFromCreateUserPayload Converts source CreateUserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromCreateUserPayload(payload *app.CreateUserPayload) *User {
	user := &User{}
	user.Email = payload.Email
	user.FirstName = payload.FirstName
	if payload.ID != nil {
		user.ID = *payload.ID
	}
	user.LastName = payload.LastName
	user.Password = payload.Password
	user.Role = payload.Role
	if payload.Validated != nil {
		user.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		user.ValidationCode = payload.ValidationCode
	}

	return user
}

// UpdateFromCreateUserPayload applies non-nil changes from CreateUserPayload to the model and saves it
func (m *UserDB) UpdateFromCreateUserPayload(ctx context.Context, payload *app.CreateUserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromcreateUserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	obj.Email = payload.Email
	obj.FirstName = payload.FirstName
	if payload.ID != nil {
		obj.ID = *payload.ID
	}
	obj.LastName = payload.LastName
	obj.Password = payload.Password
	obj.Role = payload.Role
	if payload.Validated != nil {
		obj.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		obj.ValidationCode = payload.ValidationCode
	}

	err = m.Db.Save(&obj).Error
	return err
}

// UserFromUpdateAdminuserPayload Converts source UpdateAdminuserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromUpdateAdminuserPayload(payload *app.UpdateAdminuserPayload) *User {
	user := &User{}
	if payload.Email != nil {
		user.Email = *payload.Email
	}
	if payload.FirstName != nil {
		user.FirstName = *payload.FirstName
	}
	if payload.ID != nil {
		user.ID = *payload.ID
	}
	if payload.LastName != nil {
		user.LastName = *payload.LastName
	}
	if payload.Password != nil {
		user.Password = *payload.Password
	}
	if payload.Role != nil {
		user.Role = *payload.Role
	}
	if payload.Validated != nil {
		user.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		user.ValidationCode = payload.ValidationCode
	}

	return user
}

// UpdateFromUpdateAdminuserPayload applies non-nil changes from UpdateAdminuserPayload to the model and saves it
func (m *UserDB) UpdateFromUpdateAdminuserPayload(ctx context.Context, payload *app.UpdateAdminuserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromupdateAdminuserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	if payload.Email != nil {
		obj.Email = *payload.Email
	}
	if payload.FirstName != nil {
		obj.FirstName = *payload.FirstName
	}
	if payload.ID != nil {
		obj.ID = *payload.ID
	}
	if payload.LastName != nil {
		obj.LastName = *payload.LastName
	}
	if payload.Password != nil {
		obj.Password = *payload.Password
	}
	if payload.Role != nil {
		obj.Role = *payload.Role
	}
	if payload.Validated != nil {
		obj.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		obj.ValidationCode = payload.ValidationCode
	}

	err = m.Db.Save(&obj).Error
	return err
}

// UserFromUpdateUserPayload Converts source UpdateUserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromUpdateUserPayload(payload *app.UpdateUserPayload) *User {
	user := &User{}
	if payload.Email != nil {
		user.Email = *payload.Email
	}
	if payload.FirstName != nil {
		user.FirstName = *payload.FirstName
	}
	if payload.ID != nil {
		user.ID = *payload.ID
	}
	if payload.LastName != nil {
		user.LastName = *payload.LastName
	}
	if payload.Password != nil {
		user.Password = *payload.Password
	}
	if payload.Role != nil {
		user.Role = *payload.Role
	}
	if payload.Validated != nil {
		user.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		user.ValidationCode = payload.ValidationCode
	}

	return user
}

// UpdateFromUpdateUserPayload applies non-nil changes from UpdateUserPayload to the model and saves it
func (m *UserDB) UpdateFromUpdateUserPayload(ctx context.Context, payload *app.UpdateUserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromupdateUserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	if payload.Email != nil {
		obj.Email = *payload.Email
	}
	if payload.FirstName != nil {
		obj.FirstName = *payload.FirstName
	}
	if payload.ID != nil {
		obj.ID = *payload.ID
	}
	if payload.LastName != nil {
		obj.LastName = *payload.LastName
	}
	if payload.Password != nil {
		obj.Password = *payload.Password
	}
	if payload.Role != nil {
		obj.Role = *payload.Role
	}
	if payload.Validated != nil {
		obj.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		obj.ValidationCode = payload.ValidationCode
	}

	err = m.Db.Save(&obj).Error
	return err
}
