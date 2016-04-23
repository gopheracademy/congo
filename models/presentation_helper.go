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

// ListPresentationAdmin returns an array of view: admin.
func (m *PresentationDB) ListPresentationAdmin(ctx context.Context, eventID int, speakerID int) []*app.PresentationAdmin {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "listpresentationadmin"}, time.Now())

	var native []*Presentation
	var objs []*app.PresentationAdmin
	err := m.Db.Scopes(PresentationFilterByEvent(eventID, &m.Db), PresentationFilterBySpeaker(speakerID, &m.Db)).Table(m.TableName()).Preload("Speaker").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Presentation", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PresentationToPresentationAdmin())
	}

	return objs
}

// PresentationToPresentationAdmin returns the Presentation representation of Presentation.
func (m *Presentation) PresentationToPresentationAdmin() *app.PresentationAdmin {
	presentation := &app.PresentationAdmin{}
	tmp1 := m.Speaker.SpeakerToSpeakerLink()
	presentation.Links = &app.PresentationLinks{Speaker: tmp1}
	presentation.Abstract = &m.Abstract
	presentation.Detail = m.Detail
	presentation.ID = &m.ID
	presentation.Name = m.Name

	return presentation
}

// OnePresentationAdmin returns an array of view: admin.
func (m *PresentationDB) OnePresentationAdmin(ctx context.Context, id int, eventID int, speakerID int) (*app.PresentationAdmin, error) {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "onepresentationadmin"}, time.Now())

	var native Presentation
	err := m.Db.Scopes(PresentationFilterByEvent(eventID, &m.Db), PresentationFilterBySpeaker(speakerID, &m.Db)).Table(m.TableName()).Preload("Event").Preload("Speaker").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Presentation", "error", err.Error())
		return nil, err
	}

	view := *native.PresentationToPresentationAdmin()
	return &view, err
}

// MediaType Retrieval Functions

// ListPresentation returns an array of view: default.
func (m *PresentationDB) ListPresentation(ctx context.Context, eventID int, speakerID int) []*app.Presentation {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "listpresentation"}, time.Now())

	var native []*Presentation
	var objs []*app.Presentation
	err := m.Db.Scopes(PresentationFilterByEvent(eventID, &m.Db), PresentationFilterBySpeaker(speakerID, &m.Db)).Table(m.TableName()).Preload("Speaker").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Presentation", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PresentationToPresentation())
	}

	return objs
}

// PresentationToPresentation returns the Presentation representation of Presentation.
func (m *Presentation) PresentationToPresentation() *app.Presentation {
	presentation := &app.Presentation{}
	tmp1 := m.Speaker.SpeakerToSpeakerLink()
	presentation.Links = &app.PresentationLinks{Speaker: tmp1}
	presentation.Abstract = &m.Abstract
	presentation.ID = &m.ID
	presentation.Name = m.Name

	return presentation
}

// OnePresentation returns an array of view: default.
func (m *PresentationDB) OnePresentation(ctx context.Context, id int, eventID int, speakerID int) (*app.Presentation, error) {
	defer goa.MeasureSince([]string{"goa", "db", "presentation", "onepresentation"}, time.Now())

	var native Presentation
	err := m.Db.Scopes(PresentationFilterByEvent(eventID, &m.Db), PresentationFilterBySpeaker(speakerID, &m.Db)).Table(m.TableName()).Preload("Event").Preload("Speaker").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Presentation", "error", err.Error())
		return nil, err
	}

	view := *native.PresentationToPresentation()
	return &view, err
}
