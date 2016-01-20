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

import "github.com/raphael/goa"

// ProposalPayload type
type ProposalPayload struct {
	Abstract  *string
	Detail    *string
	Firstname *string
	Title     *string
	Withdrawn *bool
}

// Validate validates the type instance.
func (ut *ProposalPayload) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 50, true, err)
		}
	}
	if ut.Abstract != nil {
		if len(*ut.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 500, false, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 100, true, err)
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 2000, false, err)
		}
	}
	if ut.Firstname != nil {
		if len(*ut.Firstname) < 2 {
			err = goa.InvalidLengthError(`response.firstname`, *ut.Firstname, len(*ut.Firstname), 2, true, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 10, true, err)
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 200, false, err)
		}
	}
	return
}

// MarshalProposalPayload validates and renders an instance of ProposalPayload into a interface{}
func MarshalProposalPayload(source *ProposalPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp3 := map[string]interface{}{
		"abstract":  source.Abstract,
		"detail":    source.Detail,
		"firstname": source.Firstname,
		"title":     source.Title,
		"withdrawn": source.Withdrawn,
	}
	target = tmp3
	return
}

// ReviewPayload type
type ReviewPayload struct {
	Comment *string
	Rating  *int
}

// Validate validates the type instance.
func (ut *ReviewPayload) Validate() (err error) {
	if ut.Comment != nil {
		if len(*ut.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 10, true, err)
		}
	}
	if ut.Comment != nil {
		if len(*ut.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 200, false, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true, err)
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false, err)
		}
	}
	return
}

// MarshalReviewPayload validates and renders an instance of ReviewPayload into a interface{}
func MarshalReviewPayload(source *ReviewPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp4 := map[string]interface{}{
		"comment": source.Comment,
		"rating":  source.Rating,
	}
	target = tmp4
	return
}

// UserPayload type
type UserPayload struct {
	Bio       *string
	City      *string
	Country   *string
	Email     *string
	Firstname *string
	Lastname  *string
	State     *string
}

// Validate validates the type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Bio != nil {
		if len(*ut.Bio) > 500 {
			err = goa.InvalidLengthError(`response.bio`, *ut.Bio, len(*ut.Bio), 500, false, err)
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalUserPayload validates and renders an instance of UserPayload into a interface{}
func MarshalUserPayload(source *UserPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp5 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"lastname":  source.Lastname,
		"state":     source.State,
	}
	target = tmp5
	return
}
