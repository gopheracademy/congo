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
	"github.com/gopheracademy/congo/gorma"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"src.xor.exchange/xor/xor/src/xor/cmd/api/app/v1"
)

// app.User model type
// Identifier: User

type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	One(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, o User) (User, error)
	Update(ctx context.Context, o User) error
	Delete(ctx context.Context, id int) error
}

func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{gorma.UserDB{Db: db}}

}

type User struct {
	gorma.User
}
type UserDB struct {
	gorma.UserDB
}

// Payload Conversion Helpers

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

// Version Conversion Helpers

func (m User) ToV1() *v1.User {
	target := v1.User{}
	copier.Copy(&target, &m)
	return &target
}
