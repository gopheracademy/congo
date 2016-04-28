//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"io"
	"time"
)

// UserCollection media type is a collection of User.
type UserCollection []*User

// DecodeUserCollection decodes the UserCollection instance encoded in r.
func DecodeUserCollection(r io.Reader, decoderFn goa.DecoderFunc) (UserCollection, error) {
	var decoded UserCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A user belonging to a tenant
type User struct {
	// Activation status
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// Email address of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// First name of user
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// API href of record
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Last name of user
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// User's role
	Role     *string `json:"role,omitempty" xml:"role,omitempty"`
	TenantID *int    `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// Validation status
	Validated *bool `json:"validated,omitempty" xml:"validated,omitempty"`
}

// DecodeUser decodes the User instance encoded in r.
func DecodeUser(r io.Reader, decoderFn goa.DecoderFunc) (*User, error) {
	var decoded User
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// Token authorization response
type Authorize struct {
	// access token
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// Time to expiration in seconds
	ExpiresIn *int `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// tenant
	Tenant *Tenant `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// type of token
	TokenType *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
	// user
	User *User `json:"user,omitempty" xml:"user,omitempty"`
}

// DecodeAuthorize decodes the Authorize instance encoded in r.
func DecodeAuthorize(r io.Reader, decoderFn goa.DecoderFunc) (*Authorize, error) {
	var decoded Authorize
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// EventCollection media type is a collection of Event.
type EventCollection []*Event

// DecodeEventCollection decodes the EventCollection instance encoded in r.
func DecodeEventCollection(r io.Reader, decoderFn goa.DecoderFunc) (EventCollection, error) {
	var decoded EventCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// An event is a specific instance of a conference, e.g. GopherCon 2016
type Event struct {
	// end_date
	EndDate *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// location
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// name
	Name          *string                `json:"name,omitempty" xml:"name,omitempty"`
	Presentations PresentationCollection `json:"presentations,omitempty" xml:"presentations,omitempty"`
	Speakers      SpeakerCollection      `json:"speakers,omitempty" xml:"speakers,omitempty"`
	// start_date
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// event URL
	URL *string `json:"url,omitempty" xml:"url,omitempty"`
}

// DecodeEvent decodes the Event instance encoded in r.
func DecodeEvent(r io.Reader, decoderFn goa.DecoderFunc) (*Event, error) {
	var decoded Event
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// PresentationCollection media type is a collection of Presentation.
type PresentationCollection []*Presentation

// DecodePresentationCollection decodes the PresentationCollection instance encoded in r.
func DecodePresentationCollection(r io.Reader, decoderFn goa.DecoderFunc) (PresentationCollection, error) {
	var decoded PresentationCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A presentation is a scheduled talk given by a speaker
type Presentation struct {
	// short description of talk
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	// detailed description of talk - not for public display
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// name of presentation
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// speaker who is giving the talk
	Speaker *Speaker `json:"speaker,omitempty" xml:"speaker,omitempty"`
}

// DecodePresentation decodes the Presentation instance encoded in r.
func DecodePresentation(r io.Reader, decoderFn goa.DecoderFunc) (*Presentation, error) {
	var decoded Presentation
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// SeriesCollection media type is a collection of Series.
type SeriesCollection []*Series

// DecodeSeriesCollection decodes the SeriesCollection instance encoded in r.
func DecodeSeriesCollection(r io.Reader, decoderFn goa.DecoderFunc) (SeriesCollection, error) {
	var decoded SeriesCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A series is a recurring event, e.g. GopherCon
type Series struct {
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// DecodeSeries decodes the Series instance encoded in r.
func DecodeSeries(r io.Reader, decoderFn goa.DecoderFunc) (*Series, error) {
	var decoded Series
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// SpeakerCollection media type is a collection of Speaker.
type SpeakerCollection []*Speaker

// DecodeSpeakerCollection decodes the SpeakerCollection instance encoded in r.
func DecodeSpeakerCollection(r io.Reader, decoderFn goa.DecoderFunc) (SpeakerCollection, error) {
	var decoded SpeakerCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// A speaker is someone giving a talk at an event
type Speaker struct {
	// speaker bio
	Bio *string `json:"bio,omitempty" xml:"bio,omitempty"`
	// email address
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// github handle
	Github *string `json:"github,omitempty" xml:"github,omitempty"`
	Href   *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// url of speaker image
	ImageURL *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// last name
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// linkedin url
	Linkedin *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	// twitter handle - no @
	Twitter *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// DecodeSpeaker decodes the Speaker instance encoded in r.
func DecodeSpeaker(r io.Reader, decoderFn goa.DecoderFunc) (*Speaker, error) {
	var decoded Speaker
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}

// TenantCollection media type is a collection of Tenant.
type TenantCollection []*Tenant

// DecodeTenantCollection decodes the TenantCollection instance encoded in r.
func DecodeTenantCollection(r io.Reader, decoderFn goa.DecoderFunc) (TenantCollection, error) {
	var decoded TenantCollection
	err := decoderFn(r).Decode(&decoded)
	return decoded, err
}

// Tenant account in Congo
type Tenant struct {
	// ID
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// DecodeTenant decodes the Tenant instance encoded in r.
func DecodeTenant(r io.Reader, decoderFn goa.DecoderFunc) (*Tenant, error) {
	var decoded Tenant
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}
