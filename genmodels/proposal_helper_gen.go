//************************************************************************//
// API "congo": Model Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package genmodels

import (
	"github.com/goadesign/goa"
	"time"
)

// v1

// MediaType Retrieval Functions
// ListProposal returns an array of view: default
func (m *ProposalDB) ListV1Proposal(ctx goa.Context) []v1.Proposal {
	now := time.Now()
	defer ctx.Info("ListProposal", "duration", time.Since(now))
	var objs []v1.Proposal
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Proposal) ProposalToV1Proposal() *v1.Proposal {
	proposal := &v1.Proposal{}
	proposal.Abstract = m.Abstract
	proposal.ID = &m.ID
	proposal.Detail = m.Detail
	proposal.Title = m.Title

	return proposal
}

// OneV1Proposal returns an array of view: default
func (m *ProposalDB) OneProposal(ctx goa.Context, id int) *v1.Proposal {
	now := time.Now()
	defer ctx.Info("OneProposal", "duration", time.Since(now))

	var native Proposal

	m.Db.Table(m.TableName()).Preload("Review").Preload("User").Where("id = ?", id).Find(&native)
	view := native.ProposalToV1Proposal()
	return view

}

// v1

// MediaType Retrieval Functions
// ListProposalLink returns an array of view: link
func (m *ProposalDB) ListV1ProposalLink(ctx goa.Context) []v1.ProposalLink {
	now := time.Now()
	defer ctx.Info("ListProposalLink", "duration", time.Since(now))
	var objs []v1.ProposalLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Proposal", "error", err.Error())
		return objs
	}

	return objs
}

func (m *Proposal) ProposalToV1ProposalLink() *v1.ProposalLink {
	proposal := &v1.ProposalLink{}
	proposal.ID = &m.ID
	proposal.Title = m.Title

	return proposal
}

// OneV1ProposalLink returns an array of view: link
func (m *ProposalDB) OneProposalLink(ctx goa.Context, id int) *v1.ProposalLink {
	now := time.Now()
	defer ctx.Info("OneProposalLink", "duration", time.Since(now))

	var native Proposal

	m.Db.Table(m.TableName()).Preload("Review").Preload("User").Where("id = ?", id).Find(&native)
	view := native.ProposalToV1ProposalLink()
	return view

}
