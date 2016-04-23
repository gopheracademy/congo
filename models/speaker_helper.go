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

// ListSpeakerAdmin returns an array of view: admin.
func (m *SpeakerDB) ListSpeakerAdmin(ctx context.Context, eventID int) []*app.SpeakerAdmin {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "listspeakeradmin"}, time.Now())

	var native []*Speaker
	var objs []*app.SpeakerAdmin
	err := m.Db.Scopes(SpeakerFilterByEvent(eventID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Speaker", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.SpeakerToSpeakerAdmin())
	}

	return objs
}

// SpeakerToSpeakerAdmin returns the Speaker representation of Speaker.
func (m *Speaker) SpeakerToSpeakerAdmin() *app.SpeakerAdmin {
	speaker := &app.SpeakerAdmin{}
	speaker.Bio = m.Bio
	speaker.FirstName = &m.FirstName
	speaker.Github = m.Github
	speaker.ID = &m.ID
	speaker.ImageURL = m.ImageURL
	speaker.LastName = &m.LastName
	speaker.Linkedin = m.Linkedin
	speaker.Twitter = m.Twitter

	return speaker
}

// OneSpeakerAdmin returns an array of view: admin.
func (m *SpeakerDB) OneSpeakerAdmin(ctx context.Context, id int, eventID int) (*app.SpeakerAdmin, error) {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "onespeakeradmin"}, time.Now())

	var native Speaker
	err := m.Db.Scopes(SpeakerFilterByEvent(eventID, &m.Db)).Table(m.TableName()).Preload("Event").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Speaker", "error", err.Error())
		return nil, err
	}

	view := *native.SpeakerToSpeakerAdmin()
	return &view, err
}

// MediaType Retrieval Functions

// ListSpeaker returns an array of view: default.
func (m *SpeakerDB) ListSpeaker(ctx context.Context, eventID int) []*app.Speaker {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "listspeaker"}, time.Now())

	var native []*Speaker
	var objs []*app.Speaker
	err := m.Db.Scopes(SpeakerFilterByEvent(eventID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Speaker", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.SpeakerToSpeaker())
	}

	return objs
}

// SpeakerToSpeaker returns the Speaker representation of Speaker.
func (m *Speaker) SpeakerToSpeaker() *app.Speaker {
	speaker := &app.Speaker{}
	speaker.Bio = m.Bio
	speaker.FirstName = &m.FirstName
	speaker.Github = m.Github
	speaker.ID = &m.ID
	speaker.ImageURL = m.ImageURL
	speaker.LastName = &m.LastName
	speaker.Linkedin = m.Linkedin
	speaker.Twitter = m.Twitter

	return speaker
}

// OneSpeaker returns an array of view: default.
func (m *SpeakerDB) OneSpeaker(ctx context.Context, id int, eventID int) (*app.Speaker, error) {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "onespeaker"}, time.Now())

	var native Speaker
	err := m.Db.Scopes(SpeakerFilterByEvent(eventID, &m.Db)).Table(m.TableName()).Preload("Event").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Speaker", "error", err.Error())
		return nil, err
	}

	view := *native.SpeakerToSpeaker()
	return &view, err
}

// MediaType Retrieval Functions

// ListSpeakerLink returns an array of view: link.
func (m *SpeakerDB) ListSpeakerLink(ctx context.Context, eventID int) []*app.SpeakerLink {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "listspeakerlink"}, time.Now())

	var native []*Speaker
	var objs []*app.SpeakerLink
	err := m.Db.Scopes(SpeakerFilterByEvent(eventID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Speaker", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.SpeakerToSpeakerLink())
	}

	return objs
}

// SpeakerToSpeakerLink returns the Speaker representation of Speaker.
func (m *Speaker) SpeakerToSpeakerLink() *app.SpeakerLink {
	speaker := &app.SpeakerLink{}
	speaker.ID = &m.ID

	return speaker
}

// OneSpeakerLink returns an array of view: link.
func (m *SpeakerDB) OneSpeakerLink(ctx context.Context, id int, eventID int) (*app.SpeakerLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "speaker", "onespeakerlink"}, time.Now())

	var native Speaker
	err := m.Db.Scopes(SpeakerFilterByEvent(eventID, &m.Db)).Table(m.TableName()).Preload("Event").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Speaker", "error", err.Error())
		return nil, err
	}

	view := *native.SpeakerToSpeakerLink()
	return &view, err
}
