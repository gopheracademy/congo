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

package review

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/jinzhu/copier"
)

func ReviewFromV1CreatePayload(ctx *v1.CreateReviewContext) Review {
	payload := ctx.Payload
	m := Review{}
	copier.Copy(&m, payload)

	return m
}

func ReviewFromV1UpdatePayload(ctx *v1.UpdateReviewContext) Review {
	payload := ctx.Payload
	m := Review{}
	copier.Copy(&m, payload)

	return m
}
