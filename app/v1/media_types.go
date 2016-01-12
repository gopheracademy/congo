//************************************************************************//
// API "congo" version v1: Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

import "github.com/raphael/goa"

// A response to a CFP
// Identifier: application/vnd.proposal+json
type Proposal struct {
	// Response abstract
	Abstract *string
	// Response detail
	Detail *string
	// API href of user
	Href *string
	// ID of user
	ID *int
	// Response title
	Title *string
}

// Proposal views
type ProposalViewEnum string

const (
	// Proposal default view
	ProposalDefaultView ProposalViewEnum = "default"
	// Proposal link view
	ProposalLinkView ProposalViewEnum = "link"
)

// LoadProposal loads raw data into an instance of Proposal
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
	if mt.Abstract != nil {
		if len(*mt.Abstract) < 50 {
			err = goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 50, true, err)
		}
	}
	if mt.Abstract != nil {
		if len(*mt.Abstract) > 500 {
			err = goa.InvalidLengthError(`response.abstract`, *mt.Abstract, len(*mt.Abstract), 500, false, err)
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) < 100 {
			err = goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 100, true, err)
		}
	}
	if mt.Detail != nil {
		if len(*mt.Detail) > 2000 {
			err = goa.InvalidLengthError(`response.detail`, *mt.Detail, len(*mt.Detail), 2000, false, err)
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) < 10 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 10, true, err)
		}
	}
	if mt.Title != nil {
		if len(*mt.Title) > 200 {
			err = goa.InvalidLengthError(`response.title`, *mt.Title, len(*mt.Title), 200, false, err)
		}
	}
	return
}

// MarshalProposal validates and renders an instance of Proposal into a interface{}
// using view "default".
func MarshalProposal(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp102 := map[string]interface{}{
		"abstract": source.Abstract,
		"detail":   source.Detail,
		"href":     source.Href,
		"id":       source.ID,
		"title":    source.Title,
	}
	target = tmp102
	return
}

// MarshalProposalLink validates and renders an instance of Proposal into a interface{}
// using view "link".
func MarshalProposalLink(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp103 := map[string]interface{}{
		"href":  source.Href,
		"id":    source.ID,
		"title": source.Title,
	}
	target = tmp103
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp104 string
			if val, ok := v.(string); ok {
				tmp104 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp104) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp104, len(tmp104), 50, true, err)
				}
				if len(tmp104) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp104, len(tmp104), 500, false, err)
				}
			}
			target.Abstract = &tmp104
		}
		if v, ok := val["detail"]; ok {
			var tmp105 string
			if val, ok := v.(string); ok {
				tmp105 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp105) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp105, len(tmp105), 100, true, err)
				}
				if len(tmp105) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp105, len(tmp105), 2000, false, err)
				}
			}
			target.Detail = &tmp105
		}
		if v, ok := val["href"]; ok {
			var tmp106 string
			if val, ok := v.(string); ok {
				tmp106 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp106
		}
		if v, ok := val["id"]; ok {
			var tmp107 int
			if f, ok := v.(float64); ok {
				tmp107 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp107
		}
		if v, ok := val["title"]; ok {
			var tmp108 string
			if val, ok := v.(string); ok {
				tmp108 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp108) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp108, len(tmp108), 10, true, err)
				}
				if len(tmp108) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp108, len(tmp108), 200, false, err)
				}
			}
			target.Title = &tmp108
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// ProposalCollection media type
// Identifier: application/vnd.proposal+json; type=collection
type ProposalCollection []*Proposal

// LoadProposalCollection loads raw data into an instance of ProposalCollection
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
	for i, tmp109 := range mt {
		var tmp110 map[string]interface{}
		tmp110, err = MarshalProposal(tmp109, err)
		res[i] = tmp110
	}
	return
}

// Validate validates the media type instance.
func (mt ProposalCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Abstract != nil {
			if len(*e.Abstract) < 50 {
				err = goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 50, true, err)
			}
		}
		if e.Abstract != nil {
			if len(*e.Abstract) > 500 {
				err = goa.InvalidLengthError(`response[*].abstract`, *e.Abstract, len(*e.Abstract), 500, false, err)
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) < 100 {
				err = goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 100, true, err)
			}
		}
		if e.Detail != nil {
			if len(*e.Detail) > 2000 {
				err = goa.InvalidLengthError(`response[*].detail`, *e.Detail, len(*e.Detail), 2000, false, err)
			}
		}
		if e.Title != nil {
			if len(*e.Title) < 10 {
				err = goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 10, true, err)
			}
		}
		if e.Title != nil {
			if len(*e.Title) > 200 {
				err = goa.InvalidLengthError(`response[*].title`, *e.Title, len(*e.Title), 200, false, err)
			}
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
		for tmp111, v := range val {
			target[tmp111], err = UnmarshalProposal(v, err)
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
	Comment *string
	// API href of user
	Href *string
	// ID of user
	ID *int
	// Rating of proposal, from 1-5
	Rating *int
}

// Review views
type ReviewViewEnum string

const (
	// Review default view
	ReviewDefaultView ReviewViewEnum = "default"
	// Review link view
	ReviewLinkView ReviewViewEnum = "link"
)

// LoadReview loads raw data into an instance of Review
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
	if mt.Comment != nil {
		if len(*mt.Comment) < 10 {
			err = goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 10, true, err)
		}
	}
	if mt.Comment != nil {
		if len(*mt.Comment) > 200 {
			err = goa.InvalidLengthError(`response.comment`, *mt.Comment, len(*mt.Comment), 200, false, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	return
}

// MarshalReview validates and renders an instance of Review into a interface{}
// using view "default".
func MarshalReview(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp112 := map[string]interface{}{
		"comment": source.Comment,
		"href":    source.Href,
		"id":      source.ID,
		"rating":  source.Rating,
	}
	target = tmp112
	return
}

// MarshalReviewLink validates and renders an instance of Review into a interface{}
// using view "link".
func MarshalReviewLink(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp113 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp113
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["comment"]; ok {
			var tmp114 string
			if val, ok := v.(string); ok {
				tmp114 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp114) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp114, len(tmp114), 10, true, err)
				}
				if len(tmp114) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp114, len(tmp114), 200, false, err)
				}
			}
			target.Comment = &tmp114
		}
		if v, ok := val["href"]; ok {
			var tmp115 string
			if val, ok := v.(string); ok {
				tmp115 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp115
		}
		if v, ok := val["id"]; ok {
			var tmp116 int
			if f, ok := v.(float64); ok {
				tmp116 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp116
		}
		if v, ok := val["rating"]; ok {
			var tmp117 int
			if f, ok := v.(float64); ok {
				tmp117 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp117 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp117, 1, true, err)
				}
				if tmp117 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp117, 5, false, err)
				}
			}
			target.Rating = &tmp117
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// ReviewCollection media type
// Identifier: application/vnd.review+json; type=collection
type ReviewCollection []*Review

// LoadReviewCollection loads raw data into an instance of ReviewCollection
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
	for i, tmp118 := range mt {
		var tmp119 map[string]interface{}
		tmp119, err = MarshalReview(tmp118, err)
		res[i] = tmp119
	}
	return
}

// Validate validates the media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Comment != nil {
			if len(*e.Comment) < 10 {
				err = goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 10, true, err)
			}
		}
		if e.Comment != nil {
			if len(*e.Comment) > 200 {
				err = goa.InvalidLengthError(`response[*].comment`, *e.Comment, len(*e.Comment), 200, false, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
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
		for tmp120, v := range val {
			target[tmp120], err = UnmarshalReview(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
