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

package models

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListUser returns an array of view: default.
func (m *UserDB) ListUser(ctx context.Context, tenantID int) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*User
	var objs []*app.User
	err := m.Db.Scopes(UserFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUser())
	}

	return objs
}

// UserToUser returns the User representation of User.
func (m *User) UserToUser() *app.User {
	user := &app.User{}
	user.Email = &m.Email
	user.FirstName = &m.FirstName
	user.Href = &m.Href
	user.ID = &m.ID
	user.LastName = &m.LastName
	user.Role = &m.Role
	user.TenantID = &m.TenantID

	return user
}

// OneUser returns an array of view: default.
func (m *UserDB) OneUser(ctx context.Context, id int, tenantID int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native User
	err := m.Db.Scopes(UserFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Preload("Tenant").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserLink returns an array of view: link.
func (m *UserDB) ListUserLink(ctx context.Context, tenantID int) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*User
	var objs []*app.UserLink
	err := m.Db.Scopes(UserFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserLink())
	}

	return objs
}

// UserToUserLink returns the User representation of User.
func (m *User) UserToUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.Email = &m.Email
	user.Href = &m.Href
	user.ID = &m.ID

	return user
}

// OneUserLink returns an array of view: link.
func (m *UserDB) OneUserLink(ctx context.Context, id int, tenantID int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native User
	err := m.Db.Scopes(UserFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Preload("Tenant").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserTenant returns an array of view: tenant.
func (m *UserDB) ListUserTenant(ctx context.Context, tenantID int) []*app.UserTenant {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listusertenant"}, time.Now())

	var native []*User
	var objs []*app.UserTenant
	err := m.Db.Scopes(UserFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserTenant())
	}

	return objs
}

// UserToUserTenant returns the User representation of User.
func (m *User) UserToUserTenant() *app.UserTenant {
	user := &app.UserTenant{}
	user.Email = &m.Email
	user.FirstName = &m.FirstName
	user.Href = &m.Href
	user.ID = &m.ID
	user.LastName = &m.LastName
	user.Role = &m.Role
	user.TenantID = &m.TenantID
	user.Validated = m.Validated

	return user
}

// OneUserTenant returns an array of view: tenant.
func (m *UserDB) OneUserTenant(ctx context.Context, id int, tenantID int) (*app.UserTenant, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneusertenant"}, time.Now())

	var native User
	err := m.Db.Scopes(UserFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Preload("Tenant").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserTenant()
	return &view, err
}
