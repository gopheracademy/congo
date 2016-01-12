//************************************************************************//
// congo: Media Helpers
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
	"github.com/gopheracademy/congo/app/v1"
	"github.com/jinzhu/copier"
)

func ProposalFromV1CreatePayload(ctx *v1.CreateProposalContext) Proposal {
	payload := ctx.Payload
	m := Proposal{}
	copier.Copy(&m, payload)

	return m
}

func ProposalFromV1UpdatePayload(ctx *v1.UpdateProposalContext) Proposal {
	payload := ctx.Payload
	m := Proposal{}
	copier.Copy(&m, payload)

	return m
}
