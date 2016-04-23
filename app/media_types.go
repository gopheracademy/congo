//************************************************************************//
// API "congo": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// Authorize media type.
//
// Identifier: application/vnd.authorize+json
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

// Validate validates the Authorize media type instance.
func (mt *Authorize) Validate() (err error) {
	if mt.User != nil {
		if mt.User.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.User.Email); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`response.user.email`, *mt.User.Email, goa.FormatEmail, err2))
			}
		}
		if mt.User.Email != nil {
			if len(*mt.User.Email) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user.email`, *mt.User.Email, len(*mt.User.Email), 2, true))
			}
		}
		if mt.User.FirstName != nil {
			if len(*mt.User.FirstName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user.first_name`, *mt.User.FirstName, len(*mt.User.FirstName), 2, true))
			}
		}
		if mt.User.LastName != nil {
			if len(*mt.User.LastName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user.last_name`, *mt.User.LastName, len(*mt.User.LastName), 2, true))
			}
		}
	}
	return err
}

// Event media type.
//
// Identifier: application/vnd.event+json
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

// EventCollection media type is a collection of Event.
//
// Identifier: application/vnd.event+json; type=collection
type EventCollection []*Event

// Login media type.
//
// Identifier: application/vnd.login+json
type Login struct {
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// PresentationAdmin media type.
//
// Identifier: application/vnd.presentation+json
type PresentationAdmin struct {
	// short description of talk
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	// detailed description of talk - not for public display
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Links to related resources
	Links *PresentationLinks `json:"links,omitempty" xml:"links,omitempty"`
	// name of presentation
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Presentation media type.
//
// Identifier: application/vnd.presentation+json
type Presentation struct {
	// short description of talk
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// Links to related resources
	Links *PresentationLinks `json:"links,omitempty" xml:"links,omitempty"`
	// name of presentation
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// PresentationLinks contains links to related resources of Presentation.
type PresentationLinks struct {
	Speaker *SpeakerLink `json:"speaker,omitempty" xml:"speaker,omitempty"`
}

// PresentationAdminCollection media type is a collection of PresentationAdmin.
//
// Identifier: application/vnd.presentation+json; type=collection
type PresentationAdminCollection []*PresentationAdmin

// PresentationCollection media type is a collection of Presentation.
//
// Identifier: application/vnd.presentation+json; type=collection
type PresentationCollection []*Presentation

// PresentationLinksArray contains links to related resources of PresentationCollection.
type PresentationLinksArray []*PresentationLinks

// Series media type.
//
// Identifier: application/vnd.series+json
type Series struct {
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// SeriesCollection media type is a collection of Series.
//
// Identifier: application/vnd.series+json; type=collection
type SeriesCollection []*Series

// SpeakerAdmin media type.
//
// Identifier: application/vnd.speaker+json
type SpeakerAdmin struct {
	// speaker bio
	Bio *string `json:"bio,omitempty" xml:"bio,omitempty"`
	// email address
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// github handle
	Github *string `json:"github,omitempty" xml:"github,omitempty"`
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

// Speaker media type.
//
// Identifier: application/vnd.speaker+json
type Speaker struct {
	// speaker bio
	Bio *string `json:"bio,omitempty" xml:"bio,omitempty"`
	// first name
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// github handle
	Github *string `json:"github,omitempty" xml:"github,omitempty"`
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

// SpeakerLink media type.
//
// Identifier: application/vnd.speaker+json
type SpeakerLink struct {
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
}

// SpeakerAdminCollection media type is a collection of SpeakerAdmin.
//
// Identifier: application/vnd.speaker+json; type=collection
type SpeakerAdminCollection []*SpeakerAdmin

// SpeakerCollection media type is a collection of Speaker.
//
// Identifier: application/vnd.speaker+json; type=collection
type SpeakerCollection []*Speaker

// SpeakerLinkCollection media type is a collection of SpeakerLink.
//
// Identifier: application/vnd.speaker+json; type=collection
type SpeakerLinkCollection []*SpeakerLink

// Tenant media type.
//
// Identifier: application/vnd.tenant+json
type Tenant struct {
	// ID
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// TenantCollection media type is a collection of Tenant.
//
// Identifier: application/vnd.tenant+json; type=collection
type TenantCollection []*Tenant

// User media type.
//
// Identifier: application/vnd.user+json
type User struct {
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
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	if mt.Email != nil {
		if len(*mt.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *mt.Email, len(*mt.Email), 2, true))
		}
	}
	if mt.FirstName != nil {
		if len(*mt.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *mt.FirstName, len(*mt.FirstName), 2, true))
		}
	}
	if mt.LastName != nil {
		if len(*mt.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *mt.LastName, len(*mt.LastName), 2, true))
		}
	}
	return err
}

// UserLink media type.
//
// Identifier: application/vnd.user+json
type UserLink struct {
	// Email address of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// API href of record
	Href *string `json:"href,omitempty" xml:"href,omitempty"`
	// ID of record
	ID *int `json:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	if mt.Email != nil {
		if len(*mt.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *mt.Email, len(*mt.Email), 2, true))
		}
	}
	return err
}

// UserTenant media type.
//
// Identifier: application/vnd.user+json
type UserTenant struct {
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

// Validate validates the UserTenant media type instance.
func (mt *UserTenant) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	if mt.Email != nil {
		if len(*mt.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *mt.Email, len(*mt.Email), 2, true))
		}
	}
	if mt.FirstName != nil {
		if len(*mt.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *mt.FirstName, len(*mt.FirstName), 2, true))
		}
	}
	if mt.LastName != nil {
		if len(*mt.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *mt.LastName, len(*mt.LastName), 2, true))
		}
	}
	return err
}

// UserCollection media type is a collection of User.
//
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2))
			}
		}
		if e.Email != nil {
			if len(*e.Email) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].email`, *e.Email, len(*e.Email), 2, true))
			}
		}
		if e.FirstName != nil {
			if len(*e.FirstName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].first_name`, *e.FirstName, len(*e.FirstName), 2, true))
			}
		}
		if e.LastName != nil {
			if len(*e.LastName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].last_name`, *e.LastName, len(*e.LastName), 2, true))
			}
		}
	}
	return err
}

// UserTenantCollection media type is a collection of UserTenant.
//
// Identifier: application/vnd.user+json; type=collection
type UserTenantCollection []*UserTenant

// Validate validates the UserTenantCollection media type instance.
func (mt UserTenantCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Email != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Email); err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFormatError(`response[*].email`, *e.Email, goa.FormatEmail, err2))
			}
		}
		if e.Email != nil {
			if len(*e.Email) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].email`, *e.Email, len(*e.Email), 2, true))
			}
		}
		if e.FirstName != nil {
			if len(*e.FirstName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].first_name`, *e.FirstName, len(*e.FirstName), 2, true))
			}
		}
		if e.LastName != nil {
			if len(*e.LastName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].last_name`, *e.LastName, len(*e.LastName), 2, true))
			}
		}
	}
	return err
}
