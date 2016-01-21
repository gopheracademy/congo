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

package proposal

import (
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/app/v1"
	"github.com/gopheracademy/congo/models/review"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// Proposal Model
//  // Stores ProposalPayload
type Proposal struct {
	ID        int             `gorm:"primary_key"` // This is the Payload Model PK field
	Detail    string          //
	Title     string          //
	Withdrawn bool            //
	UserID    int             // belongs to User
	Firstname string          //
	Reviews   []review.Review // has many Reviews
	Abstract  string          //
	UpdatedAt date.Timestamp  // timestamp
	CreatedAt date.Timestamp  // timestamp
	DeletedAt *date.Timestamp // nullable timestamp (soft delete)
}

// ProposalDB is the implementation of the storage interface for Proposal
type ProposalDB struct {
	Db gorm.DB
}

// NewProposalDB creates a new storage type
func NewProposalDB(db gorm.DB) *ProposalDB {
	return &ProposalDB{Db: db}
}

// DB returns  the underlying database
func (m *ProposalDB) DB() interface{} {
	return &m.Db
}

// Storage Interface
type ProposalStorage interface {
	DB() interface{}
	List(ctx context.Context) []Proposal
	One(ctx context.Context, id int) (Proposal, error)
	Add(ctx context.Context, proposal Proposal) (Proposal, error)
	Update(ctx context.Context, proposal Proposal) error
	Delete(ctx context.Context, id int) error
	ListByUser(ctx context.Context, user_id int) []Proposal
	OneByUser(ctx context.Context, user_id, id int) (Proposal, error)
}

// CRUD Functions
// One returns a single record by ID
func (m *ProposalDB) One(ctx context.Context, id int) (Proposal, error) {

	var obj Proposal
	err := m.Db.Find(&obj, id).Error

	return obj, err
}

// Add creates a new record
func (m *ProposalDB) Add(ctx context.Context, model Proposal) (Proposal, error) {
	err := m.Db.Create(&model).Error
	return model, err
}

// Update modifies a single record
func (m *ProposalDB) Update(ctx context.Context, model Proposal) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record
func (m *ProposalDB) Delete(ctx context.Context, id int) error {
	var obj Proposal
	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}

// Belongs To Relationships
// ProposalFilterByUser is a gorm filter for a Belongs To relationship
func ProposalFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

// ListByUser returns an array of associated User models
func (m *ProposalDB) ListByUser(ctx context.Context, parentid int) []Proposal {
	var objs []Proposal
	m.Db.Scopes(ProposalFilterByUser(parentid, &m.Db)).Find(&objs)
	return objs
}

// OneByUser returns a single associated User model
func (m *ProposalDB) OneByUser(ctx context.Context, parentid, id int) (Proposal, error) {

	var obj Proposal
	err := m.Db.Scopes(ProposalFilterByUser(parentid, &m.Db)).Find(&obj, id).Error

	return obj, err
}

// FilterProposalByUser iterates a list and returns only those with the foreign key provided
func FilterProposalByUser(parent *int, list []Proposal) []Proposal {
	var filtered []Proposal
	for _, o := range list {
		if o.UserID == int(*parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}

// Useful conversion functions
func (m *ProposalDB) ToProposalPayload() app.ProposalPayload {
	payload := app.ProposalPayload{}
	*payload.Withdrawn = m.Withdrawn
	*payload.Abstract = m.Abstract
	*payload.Detail = m.Detail
	*payload.Title = m.Title
	*payload.Firstname = m.Firstname
	return payload
}

// Convert from	Version v1 CreateProposalPayload to Proposal
func ProposalFromv1CreateProposalPayload(t v1.CreateProposalPayload) Proposal {
	proposal := Proposal{}
	*proposal.Firstname = m.Firstname
	proposal.Abstract = m.Abstract
	proposal.Detail = m.Detail
	proposal.Title = m.Title
	*proposal.Withdrawn = m.Withdrawn
	return proposal
}

// Convert from	Version v1 UpdateProposalPayload to Proposal
func ProposalFromv1UpdateProposalPayload(t v1.UpdateProposalPayload) Proposal {
	proposal := Proposal{}
	*proposal.Abstract = m.Abstract
	*proposal.Detail = m.Detail
	*proposal.Title = m.Title
	*proposal.Withdrawn = m.Withdrawn
	*proposal.Firstname = m.Firstname
	return proposal
}
