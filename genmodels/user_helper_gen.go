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
	"time"

	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// MediaType Retrieval Functions
// ListUser returns an array of view: default
func (m *UserDB) ListAppUser(ctx *goa.Context) []*app.User {
	now := time.Now()
	defer ctx.Info("ListUser", "duration", time.Since(now))
	var objs []*app.User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

func (m *User) UserToAppUser() *app.User {
	user := &app.User{}
	user.ID = &m.ID
	user.Country = m.Country
	user.Email = m.Email
	user.Lastname = m.Lastname
	user.State = m.State
	user.Bio = m.Bio
	user.City = m.City
	user.Firstname = m.Firstname

	return user
}

// OneAppUser returns an array of view: default
func (m *UserDB) OneUser(ctx *goa.Context, id int) *app.User {
	now := time.Now()
	defer ctx.Info("OneUser", "duration", time.Since(now))

	var native User

	m.Db.Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native)
	view := *native.UserToAppUser()
	return &view

}

// MediaType Retrieval Functions
// ListUserLink returns an array of view: link
func (m *UserDB) ListAppUserLink(ctx *goa.Context) []*app.UserLink {
	now := time.Now()
	defer ctx.Info("ListUserLink", "duration", time.Since(now))
	var objs []*app.UserLink
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		ctx.Error("error listing User", "error", err.Error())
		return objs
	}

	return objs
}

func (m *User) UserToAppUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.Email = m.Email
	user.ID = &m.ID

	return user
}

// OneAppUserLink returns an array of view: link
func (m *UserDB) OneUserLink(ctx *goa.Context, id int) *app.UserLink {
	now := time.Now()
	defer ctx.Info("OneUserLink", "duration", time.Since(now))

	var native User

	m.Db.Table(m.TableName()).Preload("Proposals").Preload("Reviews").Where("id = ?", id).Find(&native)
	view := *native.UserToAppUserLink()
	return &view

}
