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
	"github.com/gopheracademy/congo/app/v1"
	"github.com/gopheracademy/congo/gorma"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// app.Review model type
// Identifier: Review

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

func NewReviewDB(db gorm.DB) *ReviewDB {
	return &ReviewDB{gorma.ReviewDB{Db: db}}

}

type Review struct {
	gorma.Review
}
type ReviewDB struct {
	gorma.ReviewDB
}

// Payload Conversion Helpers

func ReviewFromV1CreatePayload(ctx *v1.CreateReviewContext) Review {
	payload := ctx.Payload
	m := Review{}
	copier.Copy(&m, payload)

	m.ProposalID = int(ctx.ProposalID)
	m.UserID = int(ctx.UserID)
	return m
}

func ReviewFromV1UpdatePayload(ctx *v1.UpdateReviewContext) Review {
	payload := ctx.Payload
	m := Review{}
	copier.Copy(&m, payload)

	m.ProposalID = int(ctx.ProposalID)
	m.UserID = int(ctx.UserID)
	return m
}

// Version Conversion Helpers

func (m Review) ToV1() *v1.Review {
	target := v1.Review{}
	copier.Copy(&target, &m)
	return &target
}
