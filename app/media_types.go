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

// Token authorization response
// Identifier: application/vnd.authorize+json
type Authorize struct {
	// access token
	AccessToken string
	// Time to expiration in seconds
	ExpiresIn int
	// type of token
	TokenType string
}

// LoadAuthorize loads raw data into an instance of Authorize running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAuthorize(raw interface{}) (res *Authorize, err error) {
	res, err = UnmarshalAuthorize(raw, err)
	return
}

// Dump produces raw data from an instance of Authorize running all the
// validations. See LoadAuthorize for the definition of raw data.
func (mt *Authorize) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalAuthorize(mt, err)
	return
}

// Validate validates the media type instance.
func (mt *Authorize) Validate() (err error) {
	return
}

// MarshalAuthorize validates and renders an instance of Authorize into a interface{}
// using view "default".
func MarshalAuthorize(source *Authorize, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp39 := map[string]interface{}{
		"access_token": source.AccessToken,
		"expires_in":   source.ExpiresIn,
		"token_type":   source.TokenType,
	}
	target = tmp39
	return
}

// UnmarshalAuthorize unmarshals and validates a raw interface{} into an instance of Authorize
func UnmarshalAuthorize(source interface{}, inErr error) (target *Authorize, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Authorize)
		if v, ok := val["access_token"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.AccessToken`, v, "string", err)
			}
			target.AccessToken = tmp40
		}
		if v, ok := val["expires_in"]; ok {
			var tmp41 int
			if f, ok := v.(float64); ok {
				tmp41 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ExpiresIn`, v, "int", err)
			}
			target.ExpiresIn = tmp41
		}
		if v, ok := val["token_type"]; ok {
			var tmp42 string
			if val, ok := v.(string); ok {
				tmp42 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.TokenType`, v, "string", err)
			}
			target.TokenType = tmp42
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// Login media type
// Identifier: application/vnd.login+json
type Login struct {
	// UUID of requesting application
	Application string
	// email
	Email string
	// password
	Password string
}

// LoadLogin loads raw data into an instance of Login running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadLogin(raw interface{}) (res *Login, err error) {
	res, err = UnmarshalLogin(raw, err)
	return
}

// Dump produces raw data from an instance of Login running all the
// validations. See LoadLogin for the definition of raw data.
func (mt *Login) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalLogin(mt, err)
	return
}

// Validate validates the media type instance.
func (mt *Login) Validate() (err error) {
	return
}

// MarshalLogin validates and renders an instance of Login into a interface{}
// using view "default".
func MarshalLogin(source *Login, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp43 := map[string]interface{}{
		"application": source.Application,
		"email":       source.Email,
		"password":    source.Password,
	}
	target = tmp43
	return
}

// UnmarshalLogin unmarshals and validates a raw interface{} into an instance of Login
func UnmarshalLogin(source interface{}, inErr error) (target *Login, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Login)
		if v, ok := val["application"]; ok {
			var tmp44 string
			if val, ok := v.(string); ok {
				tmp44 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Application`, v, "string", err)
			}
			target.Application = tmp44
		}
		if v, ok := val["email"]; ok {
			var tmp45 string
			if val, ok := v.(string); ok {
				tmp45 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			target.Email = tmp45
		}
		if v, ok := val["password"]; ok {
			var tmp46 string
			if val, ok := v.(string); ok {
				tmp46 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Password`, v, "string", err)
			}
			target.Password = tmp46
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

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
	tmp47 := map[string]interface{}{
		"abstract": source.Abstract,
		"detail":   source.Detail,
		"href":     source.Href,
		"id":       source.ID,
		"title":    source.Title,
	}
	target = tmp47
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
	tmp48 := map[string]interface{}{
		"href":  source.Href,
		"id":    source.ID,
		"title": source.Title,
	}
	target = tmp48
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp49) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp49, 50, true, err)
				}
				if len(tmp49) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp49, 500, false, err)
				}
			}
			target.Abstract = tmp49
		}
		if v, ok := val["detail"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp50) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp50, 100, true, err)
				}
				if len(tmp50) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp50, 2000, false, err)
				}
			}
			target.Detail = tmp50
		}
		if v, ok := val["href"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp51
		}
		if v, ok := val["id"]; ok {
			var tmp52 int
			if f, ok := v.(float64); ok {
				tmp52 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp52
		}
		if v, ok := val["title"]; ok {
			var tmp53 string
			if val, ok := v.(string); ok {
				tmp53 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp53) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp53, 10, true, err)
				}
				if len(tmp53) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp53, 200, false, err)
				}
			}
			target.Title = tmp53
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
	for i, tmp54 := range mt {
		var tmp55 map[string]interface{}
		tmp55, err = MarshalProposal(tmp54, err)
		res[i] = tmp55
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
		for tmp56, v := range val {
			target[tmp56], err = UnmarshalProposal(v, err)
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
	tmp57 := map[string]interface{}{
		"comment": source.Comment,
		"href":    source.Href,
		"id":      source.ID,
		"rating":  source.Rating,
	}
	target = tmp57
	return
}

// MarshalReviewLink validates and renders an instance of Review into a interface{}
// using view "link".
func MarshalReviewLink(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp58 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp58
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["comment"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp59) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp59, 10, true, err)
				}
				if len(tmp59) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp59, 200, false, err)
				}
			}
			target.Comment = tmp59
		}
		if v, ok := val["href"]; ok {
			var tmp60 string
			if val, ok := v.(string); ok {
				tmp60 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp60
		}
		if v, ok := val["id"]; ok {
			var tmp61 int
			if f, ok := v.(float64); ok {
				tmp61 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp61
		}
		if v, ok := val["rating"]; ok {
			var tmp62 int
			if f, ok := v.(float64); ok {
				tmp62 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp62 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp62, 1, true, err)
				}
				if tmp62 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp62, 5, false, err)
				}
			}
			target.Rating = tmp62
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
	for i, tmp63 := range mt {
		var tmp64 map[string]interface{}
		tmp64, err = MarshalReview(tmp63, err)
		res[i] = tmp64
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
		for tmp65, v := range val {
			target[tmp65], err = UnmarshalReview(v, err)
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
	// User Password
	Password string
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
	if len(mt.Email) < 2 {
		err = goa.InvalidLengthError(`response.email`, mt.Email, 2, true, err)
	}
	if mt.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
			err = goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2, err)
		}
	}
	if len(mt.Password) < 8 {
		err = goa.InvalidLengthError(`response.password`, mt.Password, 8, true, err)
	}
	return
}

// MarshalUser validates and renders an instance of User into a interface{}
// using view "default".
func MarshalUser(source *User, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Bio) > 500 {
		err = goa.InvalidLengthError(`.bio`, source.Bio, 500, false, err)
	}
	if len(source.Email) < 2 {
		err = goa.InvalidLengthError(`.email`, source.Email, 2, true, err)
	}
	if source.Email != "" {
		if err2 := goa.ValidateFormat(goa.FormatEmail, source.Email); err2 != nil {
			err = goa.InvalidFormatError(`.email`, source.Email, goa.FormatEmail, err2, err)
		}
	}
	tmp66 := map[string]interface{}{
		"bio":       source.Bio,
		"city":      source.City,
		"country":   source.Country,
		"email":     source.Email,
		"firstname": source.Firstname,
		"href":      source.Href,
		"id":        source.ID,
		"lastname":  source.Lastname,
		"state":     source.State,
	}
	target = tmp66
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
	tmp67 := map[string]interface{}{
		"email": source.Email,
		"href":  source.Href,
		"id":    source.ID,
	}
	target = tmp67
	return
}

// UnmarshalUser unmarshals and validates a raw interface{} into an instance of User
func UnmarshalUser(source interface{}, inErr error) (target *User, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(User)
		if v, ok := val["bio"]; ok {
			var tmp68 string
			if val, ok := v.(string); ok {
				tmp68 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp68) > 500 {
					err = goa.InvalidLengthError(`load.Bio`, tmp68, 500, false, err)
				}
			}
			target.Bio = tmp68
		}
		if v, ok := val["city"]; ok {
			var tmp69 string
			if val, ok := v.(string); ok {
				tmp69 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.City`, v, "string", err)
			}
			target.City = tmp69
		}
		if v, ok := val["country"]; ok {
			var tmp70 string
			if val, ok := v.(string); ok {
				tmp70 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = tmp70
		}
		if v, ok := val["email"]; ok {
			var tmp71 string
			if val, ok := v.(string); ok {
				tmp71 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp71) < 2 {
					err = goa.InvalidLengthError(`load.Email`, tmp71, 2, true, err)
				}
				if tmp71 != "" {
					if err2 := goa.ValidateFormat(goa.FormatEmail, tmp71); err2 != nil {
						err = goa.InvalidFormatError(`load.Email`, tmp71, goa.FormatEmail, err2, err)
					}
				}
			}
			target.Email = tmp71
		}
		if v, ok := val["firstname"]; ok {
			var tmp72 string
			if val, ok := v.(string); ok {
				tmp72 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Firstname`, v, "string", err)
			}
			target.Firstname = tmp72
		}
		if v, ok := val["href"]; ok {
			var tmp73 string
			if val, ok := v.(string); ok {
				tmp73 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp73
		}
		if v, ok := val["id"]; ok {
			var tmp74 int
			if f, ok := v.(float64); ok {
				tmp74 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp74
		}
		if v, ok := val["lastname"]; ok {
			var tmp75 string
			if val, ok := v.(string); ok {
				tmp75 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Lastname`, v, "string", err)
			}
			target.Lastname = tmp75
		}
		if v, ok := val["password"]; ok {
			var tmp76 string
			if val, ok := v.(string); ok {
				tmp76 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Password`, v, "string", err)
			}
			if err == nil {
				if len(tmp76) < 8 {
					err = goa.InvalidLengthError(`load.Password`, tmp76, 8, true, err)
				}
			}
			target.Password = tmp76
		}
		if v, ok := val["state"]; ok {
			var tmp77 string
			if val, ok := v.(string); ok {
				tmp77 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.State`, v, "string", err)
			}
			target.State = tmp77
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
	for i, tmp78 := range mt {
		var tmp79 map[string]interface{}
		tmp79, err = MarshalUser(tmp78, err)
		res[i] = tmp79
	}
	return
}

// Validate validates the media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Bio) > 500 {
			err = goa.InvalidLengthError(`response[*].bio`, e.Bio, 500, false, err)
		}
		if len(e.Email) < 2 {
			err = goa.InvalidLengthError(`response[*].email`, e.Email, 2, true, err)
		}
		if e.Email != "" {
			if err2 := goa.ValidateFormat(goa.FormatEmail, e.Email); err2 != nil {
				err = goa.InvalidFormatError(`response[*].email`, e.Email, goa.FormatEmail, err2, err)
			}
		}
		if len(e.Password) < 8 {
			err = goa.InvalidLengthError(`response[*].password`, e.Password, 8, true, err)
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
		for tmp80, v := range val {
			target[tmp80], err = UnmarshalUser(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
