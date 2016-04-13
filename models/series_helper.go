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

// ListSeries returns an array of view: default.
func (m *SeriesDB) ListSeries(ctx context.Context, tenantID int) []*app.Series {
	defer goa.MeasureSince([]string{"goa", "db", "series", "listseries"}, time.Now())

	var native []*Series
	var objs []*app.Series
	err := m.Db.Scopes(SeriesFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Series", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.SeriesToSeries())
	}

	return objs
}

// SeriesToSeries returns the Series representation of Series.
func (m *Series) SeriesToSeries() *app.Series {
	series := &app.Series{}
	series.Name = &m.Name

	return series
}

// OneSeries returns an array of view: default.
func (m *SeriesDB) OneSeries(ctx context.Context, id int, tenantID int) (*app.Series, error) {
	defer goa.MeasureSince([]string{"goa", "db", "series", "oneseries"}, time.Now())

	var native Series
	err := m.Db.Scopes(SeriesFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Preload("Tenant").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Series", "error", err.Error())
		return nil, err
	}

	view := *native.SeriesToSeries()
	return &view, err
}
