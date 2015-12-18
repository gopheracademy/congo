//************************************************************************//
// congo: Application Media Types
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

// A response to a CFP
// Identifier: application/vnd.proposal+json
type Proposal struct {
	// Response abstract
	Abstract string
	// Response detail
	Detail string
	// API href of user
	Href string
	// ID of user
	ID int
	// Response title
	Title string
}

// Proposal views
type ProposalViewEnum string

const (
	// Proposal default view
	ProposalDefaultView ProposalViewEnum = "default"
	// Proposal link view
	ProposalLinkView ProposalViewEnum = "link"
)

// LoadProposal loads raw data into an instance of Proposal running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposal(raw interface{}) (res *Proposal, err error) {
	res, err = UnmarshalProposal(raw, err)
	return
}

// Dump produces raw data from an instance of Proposal running all the
// validations. See LoadProposal for the definition of raw data.
func (mt *Proposal) Dump(view ProposalViewEnum) (res map[string]interface{}, err error) {
	if view == ProposalDefaultView {
		res, err = MarshalProposal(mt, err)
	}
	if view == ProposalLinkView {
		res, err = MarshalProposalLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Proposal) Validate() (err error) {
	if len(mt.Abstract) < 50 {
		err = goa.InvalidLengthError(`response.abstract`, mt.Abstract, 50, true, err)
	}
	if len(mt.Abstract) > 500 {
		err = goa.InvalidLengthError(`response.abstract`, mt.Abstract, 500, false, err)
	}
	if len(mt.Detail) < 100 {
		err = goa.InvalidLengthError(`response.detail`, mt.Detail, 100, true, err)
	}
	if len(mt.Detail) > 2000 {
		err = goa.InvalidLengthError(`response.detail`, mt.Detail, 2000, false, err)
	}
	if len(mt.Title) < 10 {
		err = goa.InvalidLengthError(`response.title`, mt.Title, 10, true, err)
	}
	if len(mt.Title) > 200 {
		err = goa.InvalidLengthError(`response.title`, mt.Title, 200, false, err)
	}
	return
}

// MarshalProposal validates and renders an instance of Proposal into a interface{}
// using view "default".
func MarshalProposal(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Abstract) < 50 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, 50, true, err)
	}
	if len(source.Abstract) > 500 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, 500, false, err)
	}
	if len(source.Detail) < 100 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, 100, true, err)
	}
	if len(source.Detail) > 2000 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, 2000, false, err)
	}
	if len(source.Title) < 10 {
		err = goa.InvalidLengthError(`.title`, source.Title, 10, true, err)
	}
	if len(source.Title) > 200 {
		err = goa.InvalidLengthError(`.title`, source.Title, 200, false, err)
	}
	tmp31 := map[string]interface{}{
		"abstract": source.Abstract,
		"detail":   source.Detail,
		"href":     source.Href,
		"id":       source.ID,
		"title":    source.Title,
	}
	target = tmp31
	return
}

// MarshalProposalLink validates and renders an instance of Proposal into a interface{}
// using view "link".
func MarshalProposalLink(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Title) < 10 {
		err = goa.InvalidLengthError(`.title`, source.Title, 10, true, err)
	}
	if len(source.Title) > 200 {
		err = goa.InvalidLengthError(`.title`, source.Title, 200, false, err)
	}
	tmp32 := map[string]interface{}{
		"href":  source.Href,
		"id":    source.ID,
		"title": source.Title,
	}
	target = tmp32
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp33 string
			if val, ok := v.(string); ok {
				tmp33 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp33) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp33, 50, true, err)
				}
				if len(tmp33) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp33, 500, false, err)
				}
			}
			target.Abstract = tmp33
		}
		if v, ok := val["detail"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp34) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp34, 100, true, err)
				}
				if len(tmp34) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp34, 2000, false, err)
				}
			}
			target.Detail = tmp34
		}
		if v, ok := val["href"]; ok {
			var tmp35 string
			if val, ok := v.(string); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp35
		}
		if v, ok := val["id"]; ok {
			var tmp36 int
			if f, ok := v.(float64); ok {
				tmp36 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp36
		}
		if v, ok := val["title"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp37) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp37, 10, true, err)
				}
				if len(tmp37) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp37, 200, false, err)
				}
			}
			target.Title = tmp37
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// ProposalCollection media type
// Identifier: application/vnd.proposal+json; type=collection
type ProposalCollection []*Proposal

// LoadProposalCollection loads raw data into an instance of ProposalCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadProposalCollection(raw interface{}) (res ProposalCollection, err error) {
	res, err = UnmarshalProposalCollection(raw, err)
	return
}

// Dump produces raw data from an instance of ProposalCollection running all the
// validations. See LoadProposalCollection for the definition of raw data.
func (mt ProposalCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp38 := range mt {
		var tmp39 map[string]interface{}
		tmp39, err = MarshalProposal(tmp38, err)
		res[i] = tmp39
	}
	return
}

// Validate validates the media type instance.
func (mt ProposalCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Abstract) < 50 {
			err = goa.InvalidLengthError(`response[*].abstract`, e.Abstract, 50, true, err)
		}
		if len(e.Abstract) > 500 {
			err = goa.InvalidLengthError(`response[*].abstract`, e.Abstract, 500, false, err)
		}
		if len(e.Detail) < 100 {
			err = goa.InvalidLengthError(`response[*].detail`, e.Detail, 100, true, err)
		}
		if len(e.Detail) > 2000 {
			err = goa.InvalidLengthError(`response[*].detail`, e.Detail, 2000, false, err)
		}
		if len(e.Title) < 10 {
			err = goa.InvalidLengthError(`response[*].title`, e.Title, 10, true, err)
		}
		if len(e.Title) > 200 {
			err = goa.InvalidLengthError(`response[*].title`, e.Title, 200, false, err)
		}
	}
	return
}

// MarshalProposalCollection validates and renders an instance of ProposalCollection into a interface{}
// using view "default".
func MarshalProposalCollection(source ProposalCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalProposal(res, err)
	}
	return
}

// UnmarshalProposalCollection unmarshals and validates a raw interface{} into an instance of ProposalCollection
func UnmarshalProposalCollection(source interface{}, inErr error) (target ProposalCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Proposal, len(val))
		for tmp40, v := range val {
			target[tmp40], err = UnmarshalProposal(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}

// A review is submitted by a reviewer
// Identifier: application/vnd.review+json
type Review struct {
	// Review comments
	Comment string
	// API href of user
	Href string
	// ID of user
	ID int
	// Rating of proposal, from 1-5
	Rating int
}

// Review views
type ReviewViewEnum string

const (
	// Review default view
	ReviewDefaultView ReviewViewEnum = "default"
	// Review link view
	ReviewLinkView ReviewViewEnum = "link"
)

// LoadReview loads raw data into an instance of Review running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReview(raw interface{}) (res *Review, err error) {
	res, err = UnmarshalReview(raw, err)
	return
}

// Dump produces raw data from an instance of Review running all the
// validations. See LoadReview for the definition of raw data.
func (mt *Review) Dump(view ReviewViewEnum) (res map[string]interface{}, err error) {
	if view == ReviewDefaultView {
		res, err = MarshalReview(mt, err)
	}
	if view == ReviewLinkView {
		res, err = MarshalReviewLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Review) Validate() (err error) {
	if len(mt.Comment) < 10 {
		err = goa.InvalidLengthError(`response.comment`, mt.Comment, 10, true, err)
	}
	if len(mt.Comment) > 200 {
		err = goa.InvalidLengthError(`response.comment`, mt.Comment, 200, false, err)
	}
	if mt.Rating < 1 {
		err = goa.InvalidRangeError(`response.rating`, mt.Rating, 1, true, err)
	}
	if mt.Rating > 5 {
		err = goa.InvalidRangeError(`response.rating`, mt.Rating, 5, false, err)
	}
	return
}

// MarshalReview validates and renders an instance of Review into a interface{}
// using view "default".
func MarshalReview(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Comment) < 10 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, 10, true, err)
	}
	if len(source.Comment) > 200 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, 200, false, err)
	}
	if source.Rating < 1 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 1, true, err)
	}
	if source.Rating > 5 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 5, false, err)
	}
	tmp41 := map[string]interface{}{
		"comment": source.Comment,
		"href":    source.Href,
		"id":      source.ID,
		"rating":  source.Rating,
	}
	target = tmp41
	return
}

// MarshalReviewLink validates and renders an instance of Review into a interface{}
// using view "link".
func MarshalReviewLink(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp42 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp42
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["comment"]; ok {
			var tmp43 string
			if val, ok := v.(string); ok {
				tmp43 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp43) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp43, 10, true, err)
				}
				if len(tmp43) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp43, 200, false, err)
				}
			}
			target.Comment = tmp43
		}
		if v, ok := val["href"]; ok {
			var tmp44 string
			if val, ok := v.(string); ok {
				tmp44 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp44
		}
		if v, ok := val["id"]; ok {
			var tmp45 int
			if f, ok := v.(float64); ok {
				tmp45 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp45
		}
		if v, ok := val["rating"]; ok {
			var tmp46 int
			if f, ok := v.(float64); ok {
				tmp46 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp46 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp46, 1, true, err)
				}
				if tmp46 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp46, 5, false, err)
				}
			}
			target.Rating = tmp46
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// ReviewCollection media type
// Identifier: application/vnd.review+json; type=collection
type ReviewCollection []*Review

// LoadReviewCollection loads raw data into an instance of ReviewCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadReviewCollection(raw interface{}) (res ReviewCollection, err error) {
	res, err = UnmarshalReviewCollection(raw, err)
	return
}

// Dump produces raw data from an instance of ReviewCollection running all the
// validations. See LoadReviewCollection for the definition of raw data.
func (mt ReviewCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp47 := range mt {
		var tmp48 map[string]interface{}
		tmp48, err = MarshalReview(tmp47, err)
		res[i] = tmp48
	}
	return
}

// Validate validates the media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Comment) < 10 {
			err = goa.InvalidLengthError(`response[*].comment`, e.Comment, 10, true, err)
		}
		if len(e.Comment) > 200 {
			err = goa.InvalidLengthError(`response[*].comment`, e.Comment, 200, false, err)
		}
		if e.Rating < 1 {
			err = goa.InvalidRangeError(`response[*].rating`, e.Rating, 1, true, err)
		}
		if e.Rating > 5 {
			err = goa.InvalidRangeError(`response[*].rating`, e.Rating, 5, false, err)
		}
	}
	return
}

// MarshalReviewCollection validates and renders an instance of ReviewCollection into a interface{}
// using view "default".
func MarshalReviewCollection(source ReviewCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalReview(res, err)
	}
	return
}

// UnmarshalReviewCollection unmarshals and validates a raw interface{} into an instance of ReviewCollection
func UnmarshalReviewCollection(source interface{}, inErr error) (target ReviewCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Review, len(val))
		for tmp49, v := range val {
			target[tmp49], err = UnmarshalReview(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}

// A user belonging to a tenant account
// Identifier: application/vnd.user+json
type User struct {
	// Biography of user
	Bio string
	// City of residence
	City string
	// Country of residence
	Country string
	// Email address of user
	Email string
	// First name of user
	Firstname string
	// API href of user
	Href string
	// ID of user
	ID int
	// Last name of user
	Lastname string
	// State of residence
	State string
}

// User views
type UserViewEnum string

const (
	// User default view
	UserDefaultView UserViewEnum = "default"
	// User link view
	UserLinkView UserViewEnum = "link"
)

// LoadUser loads raw data into an instance of User running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUser(raw interface{}) (res *User, err error) {
	res, err = UnmarshalUser(raw, err)
	return
}

// Dump produces raw data from an instance of User running all the
// validations. See LoadUser for the definition of raw data.
func (mt *User) Dump(view UserViewEnum) (res map[string]interface{}, err error) {
	if view == UserDefaultView {
		res, err = MarshalUser(mt, err)
	}
	if view == UserLinkView {
		res, err = MarshalUserLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *User) Validate() (err error) {
	if len(mt.Bio) > 500 {
		err = goa.InvalidLengthError(`response.bio`, mt.Bio, 500, false, err)
	}
	if len(mt.City) < 2 {
		err = goa.InvalidLengthError(`response.city`, mt.City, 2, true, err)
	}
	if len(mt.Country) < 2 {
		err = goa.InvalidLengthError(`response.country`, mt.Country, 2, true, err)
	}
	if len(mt.Email) < 2 {
		err = goa.InvalidLengthError(`response.email`, mt.Email, 2, true, err)
	}
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	if len(mt.Firstname) < 2 {
		err = goa.InvalidLengthError(`response.firstname`, mt.Firstname, 2, true, err)
	}
	if len(mt.Lastname) < 2 {
		err = goa.InvalidLengthError(`response.lastname`, mt.Lastname, 2, true, err)
	}
	if len(mt.State) < 2 {
		err = goa.InvalidLengthError(`response.state`, mt.State, 2, true, err)
	}
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
// using view "default".
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	if len(source.Firstname) < 2 {
		err = goa.InvalidLengthError(`.firstname`, source.Firstname, 2, true, err)
	}
	if len(source.Lastname) < 2 {
		err = goa.InvalidLengthError(`.lastname`, source.Lastname, 2, true, err)
	}
	tmp50 := map[string]interface{}{
		"email":     source.Email,
		"firstname": source.Firstname,
		"href":      source.Href,
		"id":        source.ID,
		"lastname":  source.Lastname,
	}
	target = tmp50
	return
}

// MarshalUserLink validates and renders an instance of User into a interface{}
// using view "link".
func MarshalUserLink(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	tmp51 := map[string]interface{}{
		"email": source.Email,
		"href":  source.Href,
		"id":    source.ID,
	}
	target = tmp51
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["bio"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp52) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp52, 500, false, err)
				}
			}
			target.Bio = tmp52
		}
		if v, ok := val["city"]; ok {
			var tmp53 string
			if val, ok := v.(string); ok {
				tmp53 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			if err == nil {
				if len(tmp53) < 2 {
					err = goa.InvalidLengthError(`load.City`, tmp53, 2, true, err)
				}
			}
			target.City = tmp53
		}
		if v, ok := val["country"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			if err == nil {
				if len(tmp54) < 2 {
					err = goa.InvalidLengthError(`load.Country`, tmp54, 2, true, err)
				}
			}
			target.Country = tmp54
		}
		if v, ok := val["email"]; ok {
			var tmp55 string
			if val, ok := v.(string); ok {
				tmp55 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp55) < 2 {
					err = goa.InvalidLengthError(`load.Email`, tmp55, 2, true, err)
				}
				if tmp55 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp55); err2 != nil {
						err = goa.InvalidFormatError(`load.Email`, tmp55, goa.FormatEmail, err2, err)
					}
				}
			}
			target.Email = tmp55
		}
		if v, ok := val["firstname"]; ok {
			var tmp56 string
			if val, ok := v.(string); ok {
				tmp56 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp56) < 2 {
					err = goa.InvalidLengthError(`load.Firstname`, tmp56, 2, true, err)
				}
			}
			target.Firstname = tmp56
		}
		if v, ok := val["href"]; ok {
			var tmp57 string
			if val, ok := v.(string); ok {
				tmp57 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp57
		}
		if v, ok := val["id"]; ok {
			var tmp58 int
			if f, ok := v.(float64); ok {
				tmp58 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp58
		}
		if v, ok := val["lastname"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			if err == nil {
				if len(tmp59) < 2 {
					err = goa.InvalidLengthError(`load.Lastname`, tmp59, 2, true, err)
				}
			}
			target.Lastname = tmp59
		}
		if v, ok := val["state"]; ok {
			var tmp60 string
			if val, ok := v.(string); ok {
				tmp60 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			if err == nil {
				if len(tmp60) < 2 {
					err = goa.InvalidLengthError(`load.State`, tmp60, 2, true, err)
				}
			}
			target.State = tmp60
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// UserCollection media type
// Identifier: application/vnd.user+json; type=collection
type UserCollection []*User

// LoadUserCollection loads raw data into an instance of UserCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadUserCollection(raw interface{}) (res UserCollection, err error) {
	res, err = UnmarshalUserCollection(raw, err)
	return
}

// Dump produces raw data from an instance of UserCollection running all the
// validations. See LoadUserCollection for the definition of raw data.
func (mt UserCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp61 := range mt {
		var tmp62 map[string]interface{}
		tmp62, err = MarshalUser(tmp61, err)
		res[i] = tmp62
	}
	return
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Bio) > 500 {
			err = goa.InvalidLengthError(`response[*].bio`, e.Bio, 500, false, err)
		}
		if len(e.City) < 2 {
			err = goa.InvalidLengthError(`response[*].city`, e.City, 2, true, err)
		}
		if len(e.Country) < 2 {
			err = goa.InvalidLengthError(`response[*].country`, e.Country, 2, true, err)
		}
		if len(e.Email) < 2 {
			err = goa.InvalidLengthError(`response[*].email`, e.Email, 2, true, err)
		}
		if e.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, e.Email, goa.FormatEmail, err2, err)
			}
		}
		if len(e.Firstname) < 2 {
			err = goa.InvalidLengthError(`response[*].firstname`, e.Firstname, 2, true, err)
		}
		if len(e.Lastname) < 2 {
			err = goa.InvalidLengthError(`response[*].lastname`, e.Lastname, 2, true, err)
		}
		if len(e.State) < 2 {
			err = goa.InvalidLengthError(`response[*].state`, e.State, 2, true, err)
		}
	}
	return
}

// MarshalUserCollection validates and renders an instance of UserCollection into a interface{}
// using view "default".
func MarshalUserCollection(source UserCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalUser(res, err)
	}
	return
}

// UnmarshalUserCollection unmarshals and validates a raw interface{} into an instance of UserCollection
func UnmarshalUserCollection(source interface{}, inErr error) (target UserCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*User, len(val))
		for tmp63, v := range val {
			target[tmp63], err = UnmarshalUser(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
