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

import "time"

// AdminUserPayload user type.
type AdminUserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	TenantID       *int    `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// EventPayload user type.
type EventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// PresentationPayload user type.
type PresentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// SeriesPayload user type.
type SeriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// SpeakerPayload user type.
type SpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// TenantPayload user type.
type TenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// UserPayload user type.
type UserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}
