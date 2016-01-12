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

package review

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// app.Review storage type
// Identifier: Review
type Review struct {
	Comment string `json:"comment,omitempty"`
	ID      int    `json:"id,omitempty" gorm:"primary_key"`
	Rating  int    `json:"rating,omitempty"`

	// Foreign Keys
	ProposalID int
	UserID     int

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ReviewStorage interface {
	DB() interface{}
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
	Db gorm.DB
}

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
	m.Db.Scopes(ReviewFilterByProposal(parentid, &m.Db)).Find(&objs)
	return objs
}

func (m *ReviewDB) OneByProposal(ctx context.Context, parentid, id int) (Review, error) {

	var obj Review

	err := m.Db.Scopes(ReviewFilterByProposal(parentid, &m.Db)).Find(&obj, id).Error

	return obj, err
}

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
	m.Db.Scopes(ReviewFilterByUser(parentid, &m.Db)).Find(&objs)
	return objs
}

func (m *ReviewDB) OneByUser(ctx context.Context, parentid, id int) (Review, error) {

	var obj Review

	err := m.Db.Scopes(ReviewFilterByUser(parentid, &m.Db)).Find(&obj, id).Error

	return obj, err
}

func NewReviewDB(db gorm.DB) *ReviewDB {

	return &ReviewDB{Db: db}

}

func (m *ReviewDB) DB() interface{} {
	return &m.Db
}

func (m *ReviewDB) List(ctx context.Context) []Review {

	var objs []Review
	m.Db.Find(&objs)
	return objs
}

func (m *ReviewDB) ListByCommentEqual(ctx context.Context, comment string) []Review {

	var objs []Review
	m.Db.Where("comment = ?", comment).Find(&objs)
	return objs
}
func (m *ReviewDB) ListByCommentLike(ctx context.Context, comment string) []Review {

	var objs []Review
	m.Db.Where("comment like ?", comment).Find(&objs)
	return objs
}

func (m *ReviewDB) ListByIdEqual(ctx context.Context, id int) []Review {

	var objs []Review
	m.Db.Where("id = ?", id).Find(&objs)
	return objs
}
func (m *ReviewDB) ListByIdLike(ctx context.Context, id int) []Review {

	var objs []Review
	m.Db.Where("id like ?", id).Find(&objs)
	return objs
}

func (m *ReviewDB) ListByRatingEqual(ctx context.Context, rating int) []Review {

	var objs []Review
	m.Db.Where("rating = ?", rating).Find(&objs)
	return objs
}
func (m *ReviewDB) ListByRatingLike(ctx context.Context, rating int) []Review {

	var objs []Review
	m.Db.Where("rating like ?", rating).Find(&objs)
	return objs
}

func (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {

	var obj Review

	err := m.Db.Find(&obj, id).Error

	return obj, err
}

func (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {
	err := m.Db.Create(&model).Error

	return model, err
}

func (m *ReviewDB) Update(ctx context.Context, model Review) error {
	obj, err := m.One(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

func (m *ReviewDB) Delete(ctx context.Context, id int) error {
	var obj Review

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		return err
	}

	return nil
}

func FilterReviewByProposal(parent int, list []Review) []Review {
	var filtered []Review
	for _, o := range list {
		if o.ProposalID == int(parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}

func FilterReviewByUser(parent int, list []Review) []Review {
	var filtered []Review
	for _, o := range list {
		if o.UserID == int(parent) {
			filtered = append(filtered, o)
		}
	}
	return filtered
}
