//************************************************************************//
// congo: Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=/home/bketelsen/src/github.com/bketelsen/congo
// --design=github.com/bketelsen/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strconv"

	"github.com/raphael/goa"
)

// CreateAccountContext provides the account create action context.
type CreateAccountContext struct {
	*goa.Context
	Payload *CreateAccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller create action.
func NewCreateAccountContext(c *goa.Context) (*CreateAccountContext, error) {
	var err error
	ctx := CreateAccountContext{Context: c}
	p, err := NewCreateAccountPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string
}

// NewCreateAccountPayload instantiates a CreateAccountPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateAccountPayload(raw interface{}) (*CreateAccountPayload, error) {
	var err error
	var p *CreateAccountPayload
	if val, ok := raw.(map[string]interface{}); ok {
		p = new(CreateAccountPayload)
		if v, ok := val["name"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			p.Name = tmp1
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, raw, "map[string]interface{}", err)
	}
	return p, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAccountContext) Created() error {
	return ctx.Respond(201, nil)
}

// DeleteAccountContext provides the account delete action context.
type DeleteAccountContext struct {
	*goa.Context
	AccountID int
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller delete action.
func NewDeleteAccountContext(c *goa.Context) (*DeleteAccountContext, error) {
	var err error
	ctx := DeleteAccountContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ShowAccountContext provides the account show action context.
type ShowAccountContext struct {
	*goa.Context
	AccountID int
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller show action.
func NewShowAccountContext(c *goa.Context) (*ShowAccountContext, error) {
	var err error
	ctx := ShowAccountContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(resp *Account, view AccountViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.account; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateAccountContext provides the account update action context.
type UpdateAccountContext struct {
	*goa.Context
	AccountID int
	Payload   *UpdateAccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the account controller update action.
func NewUpdateAccountContext(c *goa.Context) (*UpdateAccountContext, error) {
	var err error
	ctx := UpdateAccountContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewUpdateAccountPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string
}

// NewUpdateAccountPayload instantiates a UpdateAccountPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateAccountPayload(raw interface{}) (*UpdateAccountPayload, error) {
	var err error
	var p *UpdateAccountPayload
	if val, ok := raw.(map[string]interface{}); ok {
		p = new(UpdateAccountPayload)
		if v, ok := val["name"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			p.Name = tmp2
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, raw, "map[string]interface{}", err)
	}
	return p, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAccountContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// CreateSeriesContext provides the series create action context.
type CreateSeriesContext struct {
	*goa.Context
	AccountID int
	Payload   *CreateSeriesPayload
}

// NewCreateSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller create action.
func NewCreateSeriesContext(c *goa.Context) (*CreateSeriesContext, error) {
	var err error
	ctx := CreateSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	p, err := NewCreateSeriesPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateSeriesPayload is the series create action payload.
type CreateSeriesPayload struct {
	Name string
}

// NewCreateSeriesPayload instantiates a CreateSeriesPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateSeriesPayload(raw interface{}) (*CreateSeriesPayload, error) {
	var err error
	var p *CreateSeriesPayload
	if val, ok := raw.(map[string]interface{}); ok {
		p = new(CreateSeriesPayload)
		if v, ok := val["name"]; ok {
			var tmp3 string
			if val, ok := v.(string); ok {
				tmp3 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp3) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp3, 2, true, err)
				}
			}
			p.Name = tmp3
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, raw, "map[string]interface{}", err)
	}
	return p, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSeriesContext) Created() error {
	return ctx.Respond(201, nil)
}

// ListSeriesContext provides the series list action context.
type ListSeriesContext struct {
	*goa.Context
	AccountID int
}

// NewListSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller list action.
func NewListSeriesContext(c *goa.Context) (*ListSeriesContext, error) {
	var err error
	ctx := ListSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSeriesContext) OK(resp SeriesCollection, view SeriesCollectionViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.series; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowSeriesContext provides the series show action context.
type ShowSeriesContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
}

// NewShowSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller show action.
func NewShowSeriesContext(c *goa.Context) (*ShowSeriesContext, error) {
	var err error
	ctx := ShowSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSeriesContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSeriesContext) OK(resp *Series, view SeriesViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.series; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateSeriesContext provides the series update action context.
type UpdateSeriesContext struct {
	*goa.Context
	AccountID int
	SeriesID  int
	Payload   *UpdateSeriesPayload
}

// NewUpdateSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller update action.
func NewUpdateSeriesContext(c *goa.Context) (*UpdateSeriesContext, error) {
	var err error
	ctx := UpdateSeriesContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawSeriesID, ok := c.Get("seriesID")
	if ok {
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			ctx.SeriesID = int(seriesID)
		} else {
			err = goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer", err)
		}
	}
	p, err := NewUpdateSeriesPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateSeriesPayload is the series update action payload.
type UpdateSeriesPayload struct {
	Name string
}

// NewUpdateSeriesPayload instantiates a UpdateSeriesPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateSeriesPayload(raw interface{}) (*UpdateSeriesPayload, error) {
	var err error
	var p *UpdateSeriesPayload
	if val, ok := raw.(map[string]interface{}); ok {
		p = new(UpdateSeriesPayload)
		if v, ok := val["name"]; ok {
			var tmp4 string
			if val, ok := v.(string); ok {
				tmp4 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp4) < 2 {
					err = goa.InvalidLengthError(`payload.Name`, tmp4, 2, true, err)
				}
			}
			p.Name = tmp4
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, raw, "map[string]interface{}", err)
	}
	return p, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateSeriesContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSeriesContext) NotFound() error {
	return ctx.Respond(404, nil)
}
