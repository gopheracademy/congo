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

package user

import (
	"github.com/gopheracademy/congo/models/generated/user"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
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
	return &UserDB{user.UserDB{Db: db}}

}

type User struct {
	user.User
}
type UserDB struct {
	user.UserDB
}
