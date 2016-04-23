//************************************************************************//
// API "congo": Application User Types
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

// AdminUserPayload user type.
type adminUserPayload struct {
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

// Validate validates the adminUserPayload type instance.
func (ut *adminUserPayload) Validate() (err error) {
	if ut.Email != nil {
		if len(*ut.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, len(*ut.Email), 2, true))
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	if ut.Password != nil {
		if len(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, len(*ut.Password), 8, true))
		}
	}
	if ut.ValidationCode != nil {
		if len(*ut.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.validation_code`, *ut.ValidationCode, len(*ut.ValidationCode), 8, true))
		}
	}
	return
}

// Publicize creates AdminUserPayload from adminUserPayload
func (ut *adminUserPayload) Publicize() *AdminUserPayload {
	var pub AdminUserPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.Role != nil {
		pub.Role = ut.Role
	}
	if ut.TenantID != nil {
		pub.TenantID = ut.TenantID
	}
	if ut.Validated != nil {
		pub.Validated = ut.Validated
	}
	if ut.ValidationCode != nil {
		pub.ValidationCode = ut.ValidationCode
	}
	return &pub
}

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

// Validate validates the AdminUserPayload type instance.
func (ut *AdminUserPayload) Validate() (err error) {
	if ut.Email != nil {
		if len(*ut.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, len(*ut.Email), 2, true))
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	if ut.Password != nil {
		if len(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, len(*ut.Password), 8, true))
		}
	}
	if ut.ValidationCode != nil {
		if len(*ut.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.validation_code`, *ut.ValidationCode, len(*ut.ValidationCode), 8, true))
		}
	}
	return err
}

// EventPayload user type.
type eventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the eventPayload type instance.
func (ut *eventPayload) Validate() (err error) {
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true))
		}
	}
	if ut.URL != nil {
		if len(*ut.URL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.url`, *ut.URL, len(*ut.URL), 5, true))
		}
	}
	return
}

// Publicize creates EventPayload from eventPayload
func (ut *eventPayload) Publicize() *EventPayload {
	var pub EventPayload
	if ut.EndDate != nil {
		pub.EndDate = ut.EndDate
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.StartDate != nil {
		pub.StartDate = ut.StartDate
	}
	if ut.URL != nil {
		pub.URL = ut.URL
	}
	return &pub
}

// EventPayload user type.
type EventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// Validate validates the EventPayload type instance.
func (ut *EventPayload) Validate() (err error) {
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true))
		}
	}
	if ut.URL != nil {
		if len(*ut.URL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.url`, *ut.URL, len(*ut.URL), 5, true))
		}
	}
	return err
}

// PresentationPayload user type.
type presentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the presentationPayload type instance.
func (ut *presentationPayload) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 10, true))
		}
	}
	return
}

// Publicize creates PresentationPayload from presentationPayload
func (ut *presentationPayload) Publicize() *PresentationPayload {
	var pub PresentationPayload
	if ut.Abstract != nil {
		pub.Abstract = ut.Abstract
	}
	if ut.Detail != nil {
		pub.Detail = ut.Detail
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// PresentationPayload user type.
type PresentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the PresentationPayload type instance.
func (ut *PresentationPayload) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 10, true))
		}
	}
	return err
}

// SeriesPayload user type.
type seriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the seriesPayload type instance.
func (ut *seriesPayload) Validate() (err error) {
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true))
		}
	}
	return
}

// Publicize creates SeriesPayload from seriesPayload
func (ut *seriesPayload) Publicize() *SeriesPayload {
	var pub SeriesPayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// SeriesPayload user type.
type SeriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the SeriesPayload type instance.
func (ut *SeriesPayload) Validate() (err error) {
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true))
		}
	}
	return err
}

// SpeakerPayload user type.
type speakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// Validate validates the speakerPayload type instance.
func (ut *speakerPayload) Validate() (err error) {
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.ImageURL != nil {
		if len(*ut.ImageURL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.image_url`, *ut.ImageURL, len(*ut.ImageURL), 5, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	return
}

// Publicize creates SpeakerPayload from speakerPayload
func (ut *speakerPayload) Publicize() *SpeakerPayload {
	var pub SpeakerPayload
	if ut.Bio != nil {
		pub.Bio = ut.Bio
	}
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.Github != nil {
		pub.Github = ut.Github
	}
	if ut.ImageURL != nil {
		pub.ImageURL = ut.ImageURL
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	if ut.Linkedin != nil {
		pub.Linkedin = ut.Linkedin
	}
	if ut.Twitter != nil {
		pub.Twitter = ut.Twitter
	}
	return &pub
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

// Validate validates the SpeakerPayload type instance.
func (ut *SpeakerPayload) Validate() (err error) {
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.ImageURL != nil {
		if len(*ut.ImageURL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.image_url`, *ut.ImageURL, len(*ut.ImageURL), 5, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	return err
}

// TenantPayload user type.
type tenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the tenantPayload type instance.
func (ut *tenantPayload) Validate() (err error) {
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true))
		}
	}
	return
}

// Publicize creates TenantPayload from tenantPayload
func (ut *tenantPayload) Publicize() *TenantPayload {
	var pub TenantPayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// TenantPayload user type.
type TenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the TenantPayload type instance.
func (ut *TenantPayload) Validate() (err error) {
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true))
		}
	}
	return err
}

// UserPayload user type.
type userPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Email != nil {
		if len(*ut.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, len(*ut.Email), 2, true))
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	if ut.Password != nil {
		if len(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, len(*ut.Password), 8, true))
		}
	}
	if ut.ValidationCode != nil {
		if len(*ut.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.validation_code`, *ut.ValidationCode, len(*ut.ValidationCode), 8, true))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.Role != nil {
		pub.Role = ut.Role
	}
	if ut.Validated != nil {
		pub.Validated = ut.Validated
	}
	if ut.ValidationCode != nil {
		pub.ValidationCode = ut.ValidationCode
	}
	return &pub
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

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Email != nil {
		if len(*ut.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, len(*ut.Email), 2, true))
		}
	}
	if ut.FirstName != nil {
		if len(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, len(*ut.FirstName), 2, true))
		}
	}
	if ut.LastName != nil {
		if len(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, len(*ut.LastName), 2, true))
		}
	}
	if ut.Password != nil {
		if len(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, len(*ut.Password), 8, true))
		}
	}
	if ut.ValidationCode != nil {
		if len(*ut.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.validation_code`, *ut.ValidationCode, len(*ut.ValidationCode), 8, true))
		}
	}
	return err
}
