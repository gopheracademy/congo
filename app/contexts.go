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
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.account+json; charset=utf-8")
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
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.series+json; type=collection; charset=utf-8")
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
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.series+json; charset=utf-8")
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

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	*goa.Context
	AccountID int
	Payload   *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(c *goa.Context) (*CreateUserContext, error) {
	var err error
	ctx := CreateUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
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
	Email     string
	FirstName string
	LastName  string
}

// NewCreateUserPayload instantiates a CreateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateUserPayload(raw interface{}) (*CreateUserPayload, error) {
	var err error
	var p *CreateUserPayload
	if val, ok := raw.(map[string]interface{}); ok {
		p = new(CreateUserPayload)
		if v, ok := val["email"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp5) < 2 {
					err = goa.InvalidLengthError(`payload.Email`, tmp5, 2, true, err)
				}
			}
			p.Email = tmp5
		}
		if v, ok := val["first_name"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			if err == nil {
				if len(tmp6) < 2 {
					err = goa.InvalidLengthError(`payload.FirstName`, tmp6, 2, true, err)
				}
			}
			p.FirstName = tmp6
		} else {
			err = goa.MissingAttributeError(`payload`, "first_name", err)
		}
		if v, ok := val["last_name"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.LastName`, v, "string", err)
			}
			if err == nil {
				if len(tmp7) < 2 {
					err = goa.InvalidLengthError(`payload.LastName`, tmp7, 2, true, err)
				}
			}
			p.LastName = tmp7
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, raw, "map[string]interface{}", err)
	}
	return p, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	return ctx.Respond(201, nil)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	*goa.Context
	AccountID int
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(c *goa.Context) (*ListUserContext, error) {
	var err error
	ctx := ListUserContext{Context: c}
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
func (ctx *ListUserContext) OK(resp UserCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.user+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	*goa.Context
	AccountID int
	UserID    int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(c *goa.Context) (*ShowUserContext, error) {
	var err error
	ctx := ShowUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawUserID, ok := c.Get("userID")
	if ok {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
	}
	return &ctx, err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(resp *User, view UserViewEnum) error {
	r, err := resp.Dump(view)
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.congo.api.user+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	*goa.Context
	AccountID int
	UserID    int
	Payload   *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(c *goa.Context) (*UpdateUserContext, error) {
	var err error
	ctx := UpdateUserContext{Context: c}
	rawAccountID, ok := c.Get("accountID")
	if ok {
		if accountID, err2 := strconv.Atoi(rawAccountID); err2 == nil {
			ctx.AccountID = int(accountID)
		} else {
			err = goa.InvalidParamTypeError("accountID", rawAccountID, "integer", err)
		}
	}
	rawUserID, ok := c.Get("userID")
	if ok {
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			ctx.UserID = int(userID)
		} else {
			err = goa.InvalidParamTypeError("userID", rawUserID, "integer", err)
		}
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
	Email     string
	FirstName string
	LastName  string
}

// NewUpdateUserPayload instantiates a UpdateUserPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateUserPayload(raw interface{}) (*UpdateUserPayload, error) {
	var err error
	var p *UpdateUserPayload
	if val, ok := raw.(map[string]interface{}); ok {
		p = new(UpdateUserPayload)
		if v, ok := val["email"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Email`, v, "string", err)
			}
			if err == nil {
				if len(tmp8) < 2 {
					err = goa.InvalidLengthError(`payload.Email`, tmp8, 2, true, err)
				}
			}
			p.Email = tmp8
		}
		if v, ok := val["first_name"]; ok {
			var tmp9 string
			if val, ok := v.(string); ok {
				tmp9 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.FirstName`, v, "string", err)
			}
			if err == nil {
				if len(tmp9) < 2 {
					err = goa.InvalidLengthError(`payload.FirstName`, tmp9, 2, true, err)
				}
			}
			p.FirstName = tmp9
		}
		if v, ok := val["last_name"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.LastName`, v, "string", err)
			}
			if err == nil {
				if len(tmp10) < 2 {
					err = goa.InvalidLengthError(`payload.LastName`, tmp10, 2, true, err)
				}
			}
			p.LastName = tmp10
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, raw, "map[string]interface{}", err)
	}
	return p, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	return ctx.Respond(404, nil)
}
