//************************************************************************//
// API "congo" version v1: Application Contexts
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

import (
	"fmt"
	"strconv"

	"github.com/gopheracademy/congo/app"
	"github.com/raphael/goa"
)

// CreateProposalContext provides the proposal create action context.
type CreateProposalContext struct {
	*goa.Context
	UserID  int
	Version string
	Payload *CreateProposalPayload
}

// NewCreateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller create action.
func NewCreateProposalContext(c *goa.Context) (*CreateProposalContext, error) {
	var err error
	ctx := CreateProposalContext{Context: c}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewCreateProposalPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateProposalPayload is the proposal create action payload.
type CreateProposalPayload struct {
	Abstract  *string
	Detail    *string
	Firstname *string
	Title     string
	Withdrawn *bool
}

// NewCreateProposalPayload instantiates a CreateProposalPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateProposalPayload(raw interface{}) (p *CreateProposalPayload, err error) {
	p, err = UnmarshalCreateProposalPayload(raw, err)
	return
}

// UnmarshalCreateProposalPayload unmarshals and validates a raw interface{} into an instance of CreateProposalPayload
func UnmarshalCreateProposalPayload(source interface{}, inErr error) (target *CreateProposalPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateProposalPayload)
		if v, ok := val["abstract"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp49) < 50 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp49, len(tmp49), 50, true, err)
				}
				if len(tmp49) > 500 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp49, len(tmp49), 500, false, err)
				}
			}
			target.Abstract = &tmp49
		}
		if v, ok := val["detail"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp50) < 100 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp50, len(tmp50), 100, true, err)
				}
				if len(tmp50) > 2000 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp50, len(tmp50), 2000, false, err)
				}
			}
			target.Detail = &tmp50
		}
		if v, ok := val["firstname"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp51) < 2 {
					err = goa.InvalidLengthError(`payload.Firstname`, tmp51, len(tmp51), 2, true, err)
				}
			}
			target.Firstname = &tmp51
		}
		if v, ok := val["title"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp52) < 10 {
					err = goa.InvalidLengthError(`payload.Title`, tmp52, len(tmp52), 10, true, err)
				}
				if len(tmp52) > 200 {
					err = goa.InvalidLengthError(`payload.Title`, tmp52, len(tmp52), 200, false, err)
				}
			}
			target.Title = tmp52
		} else {
			err = goa.MissingAttributeError(`payload`, "title", err)
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp53 bool
			if val, ok := v.(bool); ok {
				tmp53 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp53
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateProposalContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteProposalContext provides the proposal delete action context.
type DeleteProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
}

// NewDeleteProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller delete action.
func NewDeleteProposalContext(c *goa.Context) (*DeleteProposalContext, error) {
	var err error
	ctx := DeleteProposalContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteProposalContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListProposalContext provides the proposal list action context.
type ListProposalContext struct {
	*goa.Context
	UserID  int
	Version string
}

// NewListProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller list action.
func NewListProposalContext(c *goa.Context) (*ListProposalContext, error) {
	var err error
	ctx := ListProposalContext{Context: c}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListProposalContext) OK(resp ProposalCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowProposalContext provides the proposal show action context.
type ShowProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
}

// NewShowProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller show action.
func NewShowProposalContext(c *goa.Context) (*ShowProposalContext, error) {
	var err error
	ctx := ShowProposalContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowProposalContext) OK(resp *Proposal, view ProposalViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.proposal+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateProposalContext provides the proposal update action context.
type UpdateProposalContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
	Payload    *UpdateProposalPayload
}

// NewUpdateProposalContext parses the incoming request URL and body, performs validations and creates the
// context used by the proposal controller update action.
func NewUpdateProposalContext(c *goa.Context) (*UpdateProposalContext, error) {
	var err error
	ctx := UpdateProposalContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewUpdateProposalPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateProposalPayload is the proposal update action payload.
type UpdateProposalPayload struct {
	Abstract  *string
	Detail    *string
	Firstname *string
	Title     *string
	Withdrawn *bool
}

// NewUpdateProposalPayload instantiates a UpdateProposalPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateProposalPayload(raw interface{}) (p *UpdateProposalPayload, err error) {
	p, err = UnmarshalUpdateProposalPayload(raw, err)
	return
}

// UnmarshalUpdateProposalPayload unmarshals and validates a raw interface{} into an instance of UpdateProposalPayload
func UnmarshalUpdateProposalPayload(source interface{}, inErr error) (target *UpdateProposalPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateProposalPayload)
		if v, ok := val["abstract"]; ok {
			var tmp61 string
			if val, ok := v.(string); ok {
				tmp61 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Abstract`, v, "string", err)
			}
			if err == nil {
				if len(tmp61) < 50 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp61, len(tmp61), 50, true, err)
				}
				if len(tmp61) > 500 {
					err = goa.InvalidLengthError(`payload.Abstract`, tmp61, len(tmp61), 500, false, err)
				}
			}
			target.Abstract = &tmp61
		}
		if v, ok := val["detail"]; ok {
			var tmp62 string
			if val, ok := v.(string); ok {
				tmp62 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Detail`, v, "string", err)
			}
			if err == nil {
				if len(tmp62) < 100 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp62, len(tmp62), 100, true, err)
				}
				if len(tmp62) > 2000 {
					err = goa.InvalidLengthError(`payload.Detail`, tmp62, len(tmp62), 2000, false, err)
				}
			}
			target.Detail = &tmp62
		}
		if v, ok := val["firstname"]; ok {
			var tmp63 string
			if val, ok := v.(string); ok {
				tmp63 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			if err == nil {
				if len(tmp63) < 2 {
					err = goa.InvalidLengthError(`payload.Firstname`, tmp63, len(tmp63), 2, true, err)
				}
			}
			target.Firstname = &tmp63
		}
		if v, ok := val["title"]; ok {
			var tmp64 string
			if val, ok := v.(string); ok {
				tmp64 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Title`, v, "string", err)
			}
			if err == nil {
				if len(tmp64) < 10 {
					err = goa.InvalidLengthError(`payload.Title`, tmp64, len(tmp64), 10, true, err)
				}
				if len(tmp64) > 200 {
					err = goa.InvalidLengthError(`payload.Title`, tmp64, len(tmp64), 200, false, err)
				}
			}
			target.Title = &tmp64
		}
		if v, ok := val["withdrawn"]; ok {
			var tmp65 bool
			if val, ok := v.(bool); ok {
				tmp65 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Withdrawn`, v, "bool", err)
			}
			target.Withdrawn = &tmp65
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateProposalContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateProposalContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateReviewContext provides the review create action context.
type CreateReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
	Payload    *CreateReviewPayload
}

// NewCreateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller create action.
func NewCreateReviewContext(c *goa.Context) (*CreateReviewContext, error) {
	var err error
	ctx := CreateReviewContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewCreateReviewPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateReviewPayload is the review create action payload.
type CreateReviewPayload struct {
	Comment *string
	Rating  int
}

// NewCreateReviewPayload instantiates a CreateReviewPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateReviewPayload(raw interface{}) (p *CreateReviewPayload, err error) {
	p, err = UnmarshalCreateReviewPayload(raw, err)
	return
}

// UnmarshalCreateReviewPayload unmarshals and validates a raw interface{} into an instance of CreateReviewPayload
func UnmarshalCreateReviewPayload(source interface{}, inErr error) (target *CreateReviewPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateReviewPayload)
		if v, ok := val["comment"]; ok {
			var tmp68 string
			if val, ok := v.(string); ok {
				tmp68 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp68) < 10 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp68, len(tmp68), 10, true, err)
				}
				if len(tmp68) > 200 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp68, len(tmp68), 200, false, err)
				}
			}
			target.Comment = &tmp68
		}
		if v, ok := val["rating"]; ok {
			var tmp69 int
			if f, ok := v.(float64); ok {
				tmp69 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp69 < 1 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp69, 1, true, err)
				}
				if tmp69 > 5 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp69, 5, false, err)
				}
			}
			target.Rating = tmp69
		} else {
			err = goa.MissingAttributeError(`payload`, "rating", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateReviewContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteReviewContext provides the review delete action context.
type DeleteReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Version    string
}

// NewDeleteReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller delete action.
func NewDeleteReviewContext(c *goa.Context) (*DeleteReviewContext, error) {
	var err error
	ctx := DeleteReviewContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := c.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteReviewContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListReviewContext provides the review list action context.
type ListReviewContext struct {
	*goa.Context
	ProposalID int
	UserID     int
	Version    string
}

// NewListReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller list action.
func NewListReviewContext(c *goa.Context) (*ListReviewContext, error) {
	var err error
	ctx := ListReviewContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListReviewContext) OK(resp ReviewCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.review+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowReviewContext provides the review show action context.
type ShowReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Version    string
}

// NewShowReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller show action.
func NewShowReviewContext(c *goa.Context) (*ShowReviewContext, error) {
	var err error
	ctx := ShowReviewContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := c.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowReviewContext) OK(resp *Review, view ReviewViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.review+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateReviewContext provides the review update action context.
type UpdateReviewContext struct {
	*goa.Context
	ProposalID int
	ReviewID   int
	UserID     int
	Version    string
	Payload    *UpdateReviewPayload
}

// NewUpdateReviewContext parses the incoming request URL and body, performs validations and creates the
// context used by the review controller update action.
func NewUpdateReviewContext(c *goa.Context) (*UpdateReviewContext, error) {
	var err error
	ctx := UpdateReviewContext{Context: c}
	rawProposalID := c.Get("proposalID")
	if rawProposalID != "" {
		if proposalID, err2 := strconv.Atoi(rawProposalID); err2 == nil {
			ctx.ProposalID = int(proposalID)
		} else {
			err = goa.InvalidParamTypeError("proposalID", rawProposalID, "integer", err)
		}
	}
	rawReviewID := c.Get("reviewID")
	if rawReviewID != "" {
		if reviewID, err2 := strconv.Atoi(rawReviewID); err2 == nil {
			ctx.ReviewID = int(reviewID)
		} else {
			err = goa.InvalidParamTypeError("reviewID", rawReviewID, "integer", err)
		}
	}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewUpdateReviewPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateReviewPayload is the review update action payload.
type UpdateReviewPayload struct {
	Comment *string
	Rating  *int
}

// NewUpdateReviewPayload instantiates a UpdateReviewPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateReviewPayload(raw interface{}) (p *UpdateReviewPayload, err error) {
	p, err = UnmarshalUpdateReviewPayload(raw, err)
	return
}

// UnmarshalUpdateReviewPayload unmarshals and validates a raw interface{} into an instance of UpdateReviewPayload
func UnmarshalUpdateReviewPayload(source interface{}, inErr error) (target *UpdateReviewPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateReviewPayload)
		if v, ok := val["comment"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Comment`, v, "string", err)
			}
			if err == nil {
				if len(tmp81) < 10 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp81, len(tmp81), 10, true, err)
				}
				if len(tmp81) > 200 {
					err = goa.InvalidLengthError(`payload.Comment`, tmp81, len(tmp81), 200, false, err)
				}
			}
			target.Comment = &tmp81
		}
		if v, ok := val["rating"]; ok {
			var tmp82 int
			if f, ok := v.(float64); ok {
				tmp82 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Rating`, v, "int", err)
			}
			if err == nil {
				if tmp82 < 1 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp82, 1, true, err)
				}
				if tmp82 > 5 {
					err = goa.InvalidRangeError(`payload.Rating`, tmp82, 5, false, err)
				}
			}
			target.Rating = &tmp82
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateReviewContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateReviewContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	*goa.Context
	Version string
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewCreateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Bio       *string
	City      *string
	Country   *string
	Email     *string
	Firstname string
	Lastname  *string
	Role      *string
	State     *string
}

// NewCreateUserPayload instantiates a CreateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateUserPayload(raw interface{}) (p *CreateUserPayload, err error) {
	p, err = UnmarshalCreateUserPayload(raw, err)
	return
}

// UnmarshalCreateUserPayload unmarshals and validates a raw interface{} into an instance of CreateUserPayload
func UnmarshalCreateUserPayload(source interface{}, inErr error) (target *CreateUserPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateUserPayload)
		if v, ok := val["bio"]; ok {
			var tmp83 string
			if val, ok := v.(string); ok {
				tmp83 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp83) > 500 {
					err = goa.InvalidLengthError(`payload.Bio`, tmp83, len(tmp83), 500, false, err)
				}
			}
			target.Bio = &tmp83
		}
		if v, ok := val["city"]; ok {
			var tmp84 string
			if val, ok := v.(string); ok {
				tmp84 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = &tmp84
		}
		if v, ok := val["country"]; ok {
			var tmp85 string
			if val, ok := v.(string); ok {
				tmp85 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp85
		}
		if v, ok := val["email"]; ok {
			var tmp86 string
			if val, ok := v.(string); ok {
				tmp86 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, tmp86); err2 != nil {
					err = goa.InvalidFormatError(`payload.Email`, tmp86, goa.FormatEmail, err2, err)
				}
			}
			target.Email = &tmp86
		}
		if v, ok := val["firstname"]; ok {
			var tmp87 string
			if val, ok := v.(string); ok {
				tmp87 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = tmp87
		} else {
			err = goa.MissingAttributeError(`payload`, "firstname", err)
		}
		if v, ok := val["lastname"]; ok {
			var tmp88 string
			if val, ok := v.(string); ok {
				tmp88 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp88
		}
		if v, ok := val["role"]; ok {
			var tmp89 string
			if val, ok := v.(string); ok {
				tmp89 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = &tmp89
		}
		if v, ok := val["state"]; ok {
			var tmp90 string
			if val, ok := v.(string); ok {
				tmp90 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = &tmp90
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	*goa.Context
	UserID  int
	Version string
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(c *goa.Context) (*DeleteUserContext, error) {
	var err error
	ctx := DeleteUserContext{Context: c}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	*goa.Context
	Version string
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(resp app.UserCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.user+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	*goa.Context
	UserID  int
	Version string
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(c *goa.Context) (*ShowUserContext, error) {
	var err error
	ctx := ShowUserContext{Context: c}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(resp *app.User, view app.UserViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.user+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	*goa.Context
	UserID  int
	Version string
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(c *goa.Context) (*UpdateUserContext, error) {
	var err error
	ctx := UpdateUserContext{Context: c}
	rawUserID := c.Get("userID")
	if rawUserID != "" {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	rawVersion := c.Get("version")
	if rawVersion != "" {
		ctx.Version = rawVersion
	}
	p, err := NewUpdateUserPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Bio       *string
	City      *string
	Country   *string
	Email     string
	Firstname *string
	Lastname  *string
	Role      *string
	State     *string
}

// NewUpdateUserPayload instantiates a UpdateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateUserPayload(raw interface{}) (p *UpdateUserPayload, err error) {
	p, err = UnmarshalUpdateUserPayload(raw, err)
	return
}

// UnmarshalUpdateUserPayload unmarshals and validates a raw interface{} into an instance of UpdateUserPayload
func UnmarshalUpdateUserPayload(source interface{}, inErr error) (target *UpdateUserPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateUserPayload)
		if v, ok := val["bio"]; ok {
			var tmp94 string
			if val, ok := v.(string); ok {
				tmp94 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Bio`, v, "string", err)
			}
			if err == nil {
				if len(tmp94) > 500 {
					err = goa.InvalidLengthError(`payload.Bio`, tmp94, len(tmp94), 500, false, err)
				}
			}
			target.Bio = &tmp94
		}
		if v, ok := val["city"]; ok {
			var tmp95 string
			if val, ok := v.(string); ok {
				tmp95 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.City`, v, "string", err)
			}
			target.City = &tmp95
		}
		if v, ok := val["country"]; ok {
			var tmp96 string
			if val, ok := v.(string); ok {
				tmp96 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Country`, v, "string", err)
			}
			target.Country = &tmp96
		}
		if v, ok := val["email"]; ok {
			var tmp97 string
			if val, ok := v.(string); ok {
				tmp97 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, tmp97); err2 != nil {
					err = goa.InvalidFormatError(`payload.Email`, tmp97, goa.FormatEmail, err2, err)
				}
			}
			target.Email = tmp97
		} else {
			err = goa.MissingAttributeError(`payload`, "email", err)
		}
		if v, ok := val["firstname"]; ok {
			var tmp98 string
			if val, ok := v.(string); ok {
				tmp98 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Firstname`, v, "string", err)
			}
			target.Firstname = &tmp98
		}
		if v, ok := val["lastname"]; ok {
			var tmp99 string
			if val, ok := v.(string); ok {
				tmp99 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Lastname`, v, "string", err)
			}
			target.Lastname = &tmp99
		}
		if v, ok := val["role"]; ok {
			var tmp100 string
			if val, ok := v.(string); ok {
				tmp100 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Role`, v, "string", err)
			}
			target.Role = &tmp100
		}
		if v, ok := val["state"]; ok {
			var tmp101 string
			if val, ok := v.(string); ok {
				tmp101 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.State`, v, "string", err)
			}
			target.State = &tmp101
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}
