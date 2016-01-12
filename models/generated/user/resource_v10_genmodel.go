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

package user

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/jinzhu/copier"
)

func UserFromV1CreatePayload(ctx *v1.CreateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)

	return m
}

func UserFromV1UpdatePayload(ctx *v1.UpdateUserContext) User {
	payload := ctx.Payload
	m := User{}
	copier.Copy(&m, payload)

	return m
}
