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

// ListEvent returns an array of view: default.
func (m *EventDB) ListEvent(ctx context.Context, tenantID int) []*app.Event {
	defer goa.MeasureSince([]string{"goa", "db", "event", "listevent"}, time.Now())

	var native []*Event
	var objs []*app.Event
	err := m.Db.Scopes(EventFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Event", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.EventToEvent())
	}

	return objs
}

// EventToEvent returns the Event representation of Event.
func (m *Event) EventToEvent() *app.Event {
	event := &app.Event{}
	event.EndDate = m.EndDate
	event.ID = &m.ID
	event.Name = &m.Name
	for _, k := range m.Presentations {
		event.Presentations = append(event.Presentations, k.PresentationToPresentation())
	}
	for _, k := range m.Speakers {
		event.Speakers = append(event.Speakers, k.SpeakerToSpeaker())
	}
	event.StartDate = m.StartDate
	event.URL = m.URL

	return event
}

// OneEvent returns an array of view: default.
func (m *EventDB) OneEvent(ctx context.Context, id int, tenantID int) (*app.Event, error) {
	defer goa.MeasureSince([]string{"goa", "db", "event", "oneevent"}, time.Now())

	var native Event
	err := m.Db.Scopes(EventFilterByTenant(tenantID, &m.Db)).Table(m.TableName()).Preload("Presentations").Preload("Speakers").Preload("Tenant").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Event", "error", err.Error())
		return nil, err
	}

	view := *native.EventToEvent()
	return &view, err
}
