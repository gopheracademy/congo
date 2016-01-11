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

package proposal

import (
	"github.com/gopheracademy/congo/models/generated/proposal"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// app.Proposal model type
// Identifier: Proposal

type ProposalStorage interface {
	DB() interface{}
	List(ctx context.Context) []Proposal
	One(ctx context.Context, id int) (Proposal, error)
	Add(ctx context.Context, o Proposal) (Proposal, error)
	Update(ctx context.Context, o Proposal) error
	Delete(ctx context.Context, id int) error

	ListByUser(ctx context.Context, parentid int) []Proposal
	OneByUser(ctx context.Context, parentid, id int) (Proposal, error)
}

func NewProposalDB(db gorm.DB) *ProposalDB {
	return &ProposalDB{proposal.ProposalDB{Db: db}}

}

type Proposal struct {
	proposal.Proposal
}
type ProposalDB struct {
	proposal.ProposalDB
}
