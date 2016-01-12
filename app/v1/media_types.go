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
	if len(mt.Abstract) < 50 {
		err = goa.InvalidLengthError(`response.abstract`, mt.Abstract, len(mt.Abstract), 50, true, err)
	}
	if len(mt.Abstract) > 500 {
		err = goa.InvalidLengthError(`response.abstract`, mt.Abstract, len(mt.Abstract), 500, false, err)
	}
	if len(mt.Detail) < 100 {
		err = goa.InvalidLengthError(`response.detail`, mt.Detail, len(mt.Detail), 100, true, err)
	}
	if len(mt.Detail) > 2000 {
		err = goa.InvalidLengthError(`response.detail`, mt.Detail, len(mt.Detail), 2000, false, err)
	}
	if len(mt.Title) < 10 {
		err = goa.InvalidLengthError(`response.title`, mt.Title, len(mt.Title), 10, true, err)
	}
	if len(mt.Title) > 200 {
		err = goa.InvalidLengthError(`response.title`, mt.Title, len(mt.Title), 200, false, err)
	}
	return
}

// MarshalProposal validates and renders an instance of Proposal into a interface{}
// using view "default".
func MarshalProposal(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Abstract) < 50 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, len(source.Abstract), 50, true, err)
	}
	if len(source.Abstract) > 500 {
		err = goa.InvalidLengthError(`.abstract`, source.Abstract, len(source.Abstract), 500, false, err)
	}
	if len(source.Detail) < 100 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, len(source.Detail), 100, true, err)
	}
	if len(source.Detail) > 2000 {
		err = goa.InvalidLengthError(`.detail`, source.Detail, len(source.Detail), 2000, false, err)
	}
	if len(source.Title) < 10 {
		err = goa.InvalidLengthError(`.title`, source.Title, len(source.Title), 10, true, err)
	}
	if len(source.Title) > 200 {
		err = goa.InvalidLengthError(`.title`, source.Title, len(source.Title), 200, false, err)
	}
	tmp78 := map[string]interface{}{
		"abstract": source.Abstract,
		"detail":   source.Detail,
		"href":     source.Href,
		"id":       source.ID,
		"title":    source.Title,
	}
	target = tmp78
	return
}

// MarshalProposalLink validates and renders an instance of Proposal into a interface{}
// using view "link".
func MarshalProposalLink(source *Proposal, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if len(source.Title) < 10 {
		err = goa.InvalidLengthError(`.title`, source.Title, len(source.Title), 10, true, err)
	}
	if len(source.Title) > 200 {
		err = goa.InvalidLengthError(`.title`, source.Title, len(source.Title), 200, false, err)
	}
	tmp79 := map[string]interface{}{
		"href":  source.Href,
		"id":    source.ID,
		"title": source.Title,
	}
	target = tmp79
	return
}

// UnmarshalProposal unmarshals and validates a raw interface{} into an instance of Proposal
func UnmarshalProposal(source interface{}, inErr error) (target *Proposal, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Proposal)
		if v, ok := val["abstract"]; ok {
			var tmp80 string
			if val, ok := v.(string); ok {
				tmp80 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp80) < 50 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp80, len(tmp80), 50, true, err)
				}
				if len(tmp80) > 500 {
					err = goa.InvalidLengthError(`load.Abstract`, tmp80, len(tmp80), 500, false, err)
				}
			}
			target.Abstract = tmp80
		}
		if v, ok := val["detail"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp81) < 100 {
					err = goa.InvalidLengthError(`load.Detail`, tmp81, len(tmp81), 100, true, err)
				}
				if len(tmp81) > 2000 {
					err = goa.InvalidLengthError(`load.Detail`, tmp81, len(tmp81), 2000, false, err)
				}
			}
			target.Detail = tmp81
		}
		if v, ok := val["href"]; ok {
			var tmp82 string
			if val, ok := v.(string); ok {
				tmp82 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp82
		}
		if v, ok := val["id"]; ok {
			var tmp83 int
			if f, ok := v.(float64); ok {
				tmp83 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp83
		}
		if v, ok := val["title"]; ok {
			var tmp84 string
			if val, ok := v.(string); ok {
				tmp84 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp84) < 10 {
					err = goa.InvalidLengthError(`load.Title`, tmp84, len(tmp84), 10, true, err)
				}
				if len(tmp84) > 200 {
					err = goa.InvalidLengthError(`load.Title`, tmp84, len(tmp84), 200, false, err)
				}
			}
			target.Title = tmp84
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
	for i, tmp85 := range mt {
		var tmp86 map[string]interface{}
		tmp86, err = MarshalProposal(tmp85, err)
		res[i] = tmp86
	}
	return
}

// Validate validates the media type instance.
func (mt ProposalCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Abstract) < 50 {
			err = goa.InvalidLengthError(`response[*].abstract`, e.Abstract, len(e.Abstract), 50, true, err)
		}
		if len(e.Abstract) > 500 {
			err = goa.InvalidLengthError(`response[*].abstract`, e.Abstract, len(e.Abstract), 500, false, err)
		}
		if len(e.Detail) < 100 {
			err = goa.InvalidLengthError(`response[*].detail`, e.Detail, len(e.Detail), 100, true, err)
		}
		if len(e.Detail) > 2000 {
			err = goa.InvalidLengthError(`response[*].detail`, e.Detail, len(e.Detail), 2000, false, err)
		}
		if len(e.Title) < 10 {
			err = goa.InvalidLengthError(`response[*].title`, e.Title, len(e.Title), 10, true, err)
		}
		if len(e.Title) > 200 {
			err = goa.InvalidLengthError(`response[*].title`, e.Title, len(e.Title), 200, false, err)
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
		for tmp87, v := range val {
			target[tmp87], err = UnmarshalProposal(v, err)
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
	if len(mt.Comment) < 10 {
		err = goa.InvalidLengthError(`response.comment`, mt.Comment, len(mt.Comment), 10, true, err)
	}
	if len(mt.Comment) > 200 {
		err = goa.InvalidLengthError(`response.comment`, mt.Comment, len(mt.Comment), 200, false, err)
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
		err = goa.InvalidLengthError(`.comment`, source.Comment, len(source.Comment), 10, true, err)
	}
	if len(source.Comment) > 200 {
		err = goa.InvalidLengthError(`.comment`, source.Comment, len(source.Comment), 200, false, err)
	}
	if source.Rating < 1 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 1, true, err)
	}
	if source.Rating > 5 {
		err = goa.InvalidRangeError(`.rating`, source.Rating, 5, false, err)
	}
	tmp88 := map[string]interface{}{
		"comment": source.Comment,
		"href":    source.Href,
		"id":      source.ID,
		"rating":  source.Rating,
	}
	target = tmp88
	return
}

// MarshalReviewLink validates and renders an instance of Review into a interface{}
// using view "link".
func MarshalReviewLink(source *Review, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp89 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
	}
	target = tmp89
	return
}

// UnmarshalReview unmarshals and validates a raw interface{} into an instance of Review
func UnmarshalReview(source interface{}, inErr error) (target *Review, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Review)
		if v, ok := val["comment"]; ok {
			var tmp90 string
			if val, ok := v.(string); ok {
				tmp90 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp90) < 10 {
					err = goa.InvalidLengthError(`load.Comment`, tmp90, len(tmp90), 10, true, err)
				}
				if len(tmp90) > 200 {
					err = goa.InvalidLengthError(`load.Comment`, tmp90, len(tmp90), 200, false, err)
				}
			}
			target.Comment = tmp90
		}
		if v, ok := val["href"]; ok {
			var tmp91 string
			if val, ok := v.(string); ok {
				tmp91 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp91
		}
		if v, ok := val["id"]; ok {
			var tmp92 int
			if f, ok := v.(float64); ok {
				tmp92 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp92
		}
		if v, ok := val["rating"]; ok {
			var tmp93 int
			if f, ok := v.(float64); ok {
				tmp93 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp93 < 1 {
					err = goa.InvalidRangeError(`load.Rating`, tmp93, 1, true, err)
				}
				if tmp93 > 5 {
					err = goa.InvalidRangeError(`load.Rating`, tmp93, 5, false, err)
				}
			}
			target.Rating = tmp93
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
	for i, tmp94 := range mt {
		var tmp95 map[string]interface{}
		tmp95, err = MarshalReview(tmp94, err)
		res[i] = tmp95
	}
	return
}

// Validate validates the media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if len(e.Comment) < 10 {
			err = goa.InvalidLengthError(`response[*].comment`, e.Comment, len(e.Comment), 10, true, err)
		}
		if len(e.Comment) > 200 {
			err = goa.InvalidLengthError(`response[*].comment`, e.Comment, len(e.Comment), 200, false, err)
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
		for tmp96, v := range val {
			target[tmp96], err = UnmarshalReview(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
