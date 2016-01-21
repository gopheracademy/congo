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

package user

import (
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/app/v1"
	"github.com/gopheracademy/congo/models/proposal"
	"github.com/gopheracademy/congo/models/review"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// User Model
//  // Stores UserPayload
type User struct {
	ID        int                 `gorm:"primary_key"` // This is the User Model PK field
	Lastname  string              //
	Firstname string              //
	Bio       string              //
	Reviews   []review.Review     // has many Reviews
	Proposals []proposal.Proposal // has many Proposals
	City      string              //
	Country   string              //
	State     string              //
	Email     string              //
	UpdatedAt date.Timestamp      // timestamp
	DeletedAt *date.Timestamp     // nullable timestamp (soft delete)
	CreatedAt date.Timestamp      // timestamp
}

// UserDB is the implementation of the storage interface for User
type UserDB struct {
	Db gorm.DB
}

// NewUserDB creates a new storage type
func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns  the underlying database
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// Storage Interface
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	One(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id int) error
}

// CRUD Functions
// One returns a single record by ID
func (m *UserDB) One(ctx context.Context, id int) (User, error) {

	var obj User
	err := m.Db.Find(&obj, id).Error

	return obj, err
}

// Add creates a new record
func (m *UserDB) Add(ctx context.Context, model User) (User, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record
func (m *UserDB) Update(ctx context.Context, model User) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record
func (m *UserDB) Delete(ctx context.Context, id int) error {
	var obj User
	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}

// Useful conversion functions
func (m *UserDB) ToUserPayload() app.UserPayload {
	payload := app.UserPayload{}
	*payload.City = m.City
	*payload.Country = m.Country
	*payload.Firstname = m.Firstname
	*payload.State = m.State
	*payload.Bio = m.Bio
	*payload.Email = m.Email
	*payload.Lastname = m.Lastname
	return payload
}

// Convert from	Version v1 CreateUserPayload to User
func UserFromv1CreateUserPayload(t v1.CreateUserPayload) User {
	user := User{}
	user.Email = m.Email
	user.Lastname = m.Lastname
	*user.Bio = m.Bio
	*user.Country = m.Country
	user.Firstname = m.Firstname
	*user.State = m.State
	*user.City = m.City
	return user
}

// Convert from	Version v1 UpdateUserPayload to User
func UserFromv1UpdateUserPayload(t v1.UpdateUserPayload) User {
	user := User{}
	*user.Bio = m.Bio
	user.Email = m.Email
	*user.Lastname = m.Lastname
	*user.City = m.City
	*user.Country = m.Country
	*user.Firstname = m.Firstname
	*user.State = m.State
	return user
}
