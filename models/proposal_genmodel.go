//************************************************************************//
// congo: Models
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
	"time"

	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// app.ProposalModel storage type
// Identifier:
type Proposal struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	UserID    int
	Reviews   []Review
	Abstract  string `json:"abstract,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Title     string `json:"title,omitempty"`
	Withdrawn bool   `json:"withdrawn,omitempty"`
}

func ProposalFromCreatePayload(ctx *app.CreateProposalContext) Proposal {
	payload := ctx.Payload
	m := Proposal{}
	copier.Copy(&m, payload)

	m.UserID = int(ctx.UserID)
	return m
}

func ProposalFromUpdatePayload(ctx *app.UpdateProposalContext) Proposal {
	payload := ctx.Payload
	m := Proposal{}
	copier.Copy(&m, payload)
	return m
}

func (m Proposal) ToApp() *app.Proposal {
	target := app.Proposal{}
	copier.Copy(&target, &m)
	return &target
}

type ProposalStorage interface {
	List(ctx context.Context) []Proposal
	One(ctx context.Context, id int) (Proposal, error)
	Add(ctx context.Context, o Proposal) (Proposal, error)
	Update(ctx context.Context, o Proposal) error
	Delete(ctx context.Context, id int) error
}

type ProposalDB struct {
	DB gorm.DB
}

// would prefer to just pass a context in here, but they're all different, so can't
func ProposalFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func NewProposalDB(db gorm.DB) *ProposalDB {
	return &ProposalDB{DB: db}
}

func (m *ProposalDB) List(ctx context.Context) []Proposal {

	var objs []Proposal
	m.DB.Find(&objs)
	return objs
}

func (m *ProposalDB) One(ctx context.Context, id int) (Proposal, error) {

	var obj Proposal

	err := m.DB.Find(&obj, id).Error
	return obj, err
}

func (m *ProposalDB) Add(ctx context.Context, model Proposal) (Proposal, error) {
	err := m.DB.Create(&model).Error
	return model, err
}

func (m *ProposalDB) Update(ctx context.Context, model Proposal) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(model).Error
	return err
}

func (m *ProposalDB) Delete(ctx context.Context, id int) error {
	var obj Proposal
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		return err
	}
	return nil
}

func FilterProposalByUser(parent int, list []Proposal) []Proposal {
	filtered := make([]Proposal, 0)
	for _, o := range list {
		if o.UserID == int(parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}
