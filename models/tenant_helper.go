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

// ListTenant returns an array of view: default.
func (m *TenantDB) ListTenant(ctx context.Context) []*app.Tenant {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "listtenant"}, time.Now())

	var native []*Tenant
	var objs []*app.Tenant
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Tenant", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.TenantToTenant())
	}

	return objs
}

// TenantToTenant returns the Tenant representation of Tenant.
func (m *Tenant) TenantToTenant() *app.Tenant {
	tenant := &app.Tenant{}
	tenant.Name = &m.Name

	return tenant
}

// OneTenant returns an array of view: default.
func (m *TenantDB) OneTenant(ctx context.Context, id int) (*app.Tenant, error) {
	defer goa.MeasureSince([]string{"goa", "db", "tenant", "onetenant"}, time.Now())

	var native Tenant
	err := m.Db.Scopes().Table(m.TableName()).Preload("Users").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Tenant", "error", err.Error())
		return nil, err
	}

	view := *native.TenantToTenant()
	return &view, err
}
