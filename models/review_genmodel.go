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

// app.ReviewModel storage type
// Identifier:
type Review struct {
	ID      int    `json:"ID,omitempty" gorm:"primary_key"`
	Comment string `json:"comment,omitempty"`
	Rating  int    `json:"rating,omitempty"`

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// Foreign Keys
	ProposalID int
	UserID     int
}

func ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {
	payload := ctx.Payload
	m := Review{}
	copier.Copy(&m, payload)

	m.ProposalID = int(ctx.ProposalID)
	m.UserID = int(ctx.UserID)
	return m
}

func ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {
	payload := ctx.Payload
	m := Review{}
	copier.Copy(&m, payload)
	return m
}

func (m Review) ToApp() *app.Review {
	target := app.Review{}
	copier.Copy(&target, &m)
	return &target
}

type ReviewStorage interface {
	List(ctx context.Context) []Review
	One(ctx context.Context, id int) (Review, error)
	Add(ctx context.Context, o Review) (Review, error)
	Update(ctx context.Context, o Review) error
	Delete(ctx context.Context, id int) error

	ListByProposal(ctx context.Context, parentid int) []Review
	OneByProposal(ctx context.Context, parentid, id int) (Review, error)

	ListByUser(ctx context.Context, parentid int) []Review
	OneByUser(ctx context.Context, parentid, id int) (Review, error)
}

type ReviewDB struct {
	DB gorm.DB
}

// would prefer to just pass a context in here, but they're all different, so can't
func ReviewFilterByProposal(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if parentid > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("proposal_id = ?", parentid)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func (m *ReviewDB) ListByProposal(ctx context.Context, parentid int) []Review {

	var objs []Review
	m.DB.Scopes(ReviewFilterByProposal(parentid, &m.DB)).Find(&objs)
	return objs
}

func (m *ReviewDB) OneByProposal(ctx context.Context, parentid, id int) (Review, error) {

	var obj Review

	err := m.DB.Scopes(ReviewFilterByProposal(parentid, &m.DB)).Find(&obj, id).Error

	return obj, err
}

// would prefer to just pass a context in here, but they're all different, so can't
func ReviewFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
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

func (m *ReviewDB) ListByUser(ctx context.Context, parentid int) []Review {

	var objs []Review
	m.DB.Scopes(ReviewFilterByUser(parentid, &m.DB)).Find(&objs)
	return objs
}

func (m *ReviewDB) OneByUser(ctx context.Context, parentid, id int) (Review, error) {

	var obj Review

	err := m.DB.Scopes(ReviewFilterByUser(parentid, &m.DB)).Find(&obj, id).Error

	return obj, err
}

func NewReviewDB(db gorm.DB) *ReviewDB {

	return &ReviewDB{DB: db}

}

func (m *ReviewDB) List(ctx context.Context) []Review {

	var objs []Review
	m.DB.Find(&objs)
	return objs
}

func (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {

	var obj Review

	err := m.DB.Find(&obj, id).Error

	return obj, err
}

func (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {
	err := m.DB.Create(&model).Error

	return model, err
}

func (m *ReviewDB) Update(ctx context.Context, model Review) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.DB.Model(&obj).Updates(model).Error

	return err
}

func (m *ReviewDB) Delete(ctx context.Context, id int) error {
	var obj Review
	err := m.DB.Delete(&obj, id).Error
	if err != nil {
		return err
	}

	return nil
}

func FilterReviewByProposal(parent int, list []Review) []Review {
	filtered := make([]Review, 0)
	for _, o := range list {
		if o.ProposalID == int(parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}

func FilterReviewByUser(parent int, list []Review) []Review {
	filtered := make([]Review, 0)
	for _, o := range list {
		if o.UserID == int(parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}
