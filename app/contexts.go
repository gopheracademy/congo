//************************************************************************//
// API "congo": Application Contexts
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
	"golang.org/x/net/context"
	"strconv"
	"time"
)

// CreateAdminuserContext provides the adminuser create action context.
type CreateAdminuserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *CreateAdminuserPayload
}

// NewCreateAdminuserContext parses the incoming request URL and body, performs validations and creates the
// context used by the adminuser controller create action.
func NewCreateAdminuserContext(ctx context.Context, service *goa.Service) (*CreateAdminuserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateAdminuserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// createAdminuserPayload is the adminuser create action payload.
type createAdminuserPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *createAdminuserPayload) Validate() (err error) {
	if payload.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.Role == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "role"))
	}

	if payload.Email != nil {
		if len(*payload.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, *payload.Email, len(*payload.Email), 2, true))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	if payload.Password != nil {
		if len(*payload.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, len(*payload.Password), 8, true))
		}
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// Publicize creates CreateAdminuserPayload from createAdminuserPayload
func (payload *createAdminuserPayload) Publicize() *CreateAdminuserPayload {
	var pub CreateAdminuserPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.FirstName != nil {
		pub.FirstName = *payload.FirstName
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.LastName != nil {
		pub.LastName = *payload.LastName
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.Role != nil {
		pub.Role = *payload.Role
	}
	if payload.TenantID != nil {
		pub.TenantID = payload.TenantID
	}
	if payload.Validated != nil {
		pub.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		pub.ValidationCode = payload.ValidationCode
	}
	return &pub
}

// CreateAdminuserPayload is the adminuser create action payload.
type CreateAdminuserPayload struct {
	Email          string  `json:"email" xml:"email"`
	FirstName      string  `json:"first_name" xml:"first_name"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       string  `json:"last_name" xml:"last_name"`
	Password       string  `json:"password" xml:"password"`
	Role           string  `json:"role" xml:"role"`
	TenantID       *int    `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAdminuserPayload) Validate() (err error) {
	if payload.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "role"))
	}

	if len(payload.Email) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, payload.Email, len(payload.Email), 2, true))
	}
	if len(payload.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, payload.FirstName, len(payload.FirstName), 2, true))
	}
	if len(payload.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, payload.LastName, len(payload.LastName), 2, true))
	}
	if len(payload.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, payload.Password, len(payload.Password), 8, true))
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAdminuserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteAdminuserContext provides the adminuser delete action context.
type DeleteAdminuserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  int
}

// NewDeleteAdminuserContext parses the incoming request URL and body, performs validations and creates the
// context used by the adminuser controller delete action.
func NewDeleteAdminuserContext(ctx context.Context, service *goa.Service) (*DeleteAdminuserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteAdminuserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAdminuserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAdminuserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListAdminuserContext provides the adminuser list action context.
type ListAdminuserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewListAdminuserContext parses the incoming request URL and body, performs validations and creates the
// context used by the adminuser controller list action.
func NewListAdminuserContext(ctx context.Context, service *goa.Service) (*ListAdminuserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListAdminuserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAdminuserContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OKTenant sends a HTTP response with status code 200.
func (ctx *ListAdminuserContext) OKTenant(r UserTenantCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListAdminuserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowAdminuserContext provides the adminuser show action context.
type ShowAdminuserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  int
}

// NewShowAdminuserContext parses the incoming request URL and body, performs validations and creates the
// context used by the adminuser controller show action.
func NewShowAdminuserContext(ctx context.Context, service *goa.Service) (*ShowAdminuserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowAdminuserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAdminuserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OKTenant sends a HTTP response with status code 200.
func (ctx *ShowAdminuserContext) OKTenant(r *UserTenant) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAdminuserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateAdminuserContext provides the adminuser update action context.
type UpdateAdminuserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  int
	Payload *UpdateAdminuserPayload
}

// NewUpdateAdminuserContext parses the incoming request URL and body, performs validations and creates the
// context used by the adminuser controller update action.
func NewUpdateAdminuserContext(ctx context.Context, service *goa.Service) (*UpdateAdminuserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateAdminuserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// updateAdminuserPayload is the adminuser update action payload.
type updateAdminuserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateAdminuserPayload) Validate() (err error) {
	if payload.Email != nil {
		if len(*payload.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, *payload.Email, len(*payload.Email), 2, true))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	if payload.Password != nil {
		if len(*payload.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, len(*payload.Password), 8, true))
		}
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// Publicize creates UpdateAdminuserPayload from updateAdminuserPayload
func (payload *updateAdminuserPayload) Publicize() *UpdateAdminuserPayload {
	var pub UpdateAdminuserPayload
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.FirstName != nil {
		pub.FirstName = payload.FirstName
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.LastName != nil {
		pub.LastName = payload.LastName
	}
	if payload.Password != nil {
		pub.Password = payload.Password
	}
	if payload.Role != nil {
		pub.Role = payload.Role
	}
	if payload.Validated != nil {
		pub.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		pub.ValidationCode = payload.ValidationCode
	}
	return &pub
}

// UpdateAdminuserPayload is the adminuser update action payload.
type UpdateAdminuserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAdminuserPayload) Validate() (err error) {
	if payload.Email != nil {
		if len(*payload.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, *payload.Email, len(*payload.Email), 2, true))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	if payload.Password != nil {
		if len(*payload.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, len(*payload.Password), 8, true))
		}
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAdminuserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAdminuserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// RefreshAuthContext provides the auth refresh action context.
type RefreshAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(ctx context.Context, service *goa.Service) (*RefreshAuthContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := RefreshAuthContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *RefreshAuthContext) Created(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.Service.Send(ctx.Context, 201, r)
}

// TokenAuthContext provides the auth token action context.
type TokenAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(ctx context.Context, service *goa.Service) (*TokenAuthContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := TokenAuthContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *TokenAuthContext) Created(r *Authorize) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.Service.Send(ctx.Context, 201, r)
}

// CreateEventContext provides the event create action context.
type CreateEventContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	Payload  *CreateEventPayload
}

// NewCreateEventContext parses the incoming request URL and body, performs validations and creates the
// context used by the event controller create action.
func NewCreateEventContext(ctx context.Context, service *goa.Service) (*CreateEventContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateEventContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// createEventPayload is the event create action payload.
type createEventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createEventPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.URL != nil {
		if len(*payload.URL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.url`, *payload.URL, len(*payload.URL), 5, true))
		}
	}
	return err
}

// Publicize creates CreateEventPayload from createEventPayload
func (payload *createEventPayload) Publicize() *CreateEventPayload {
	var pub CreateEventPayload
	if payload.EndDate != nil {
		pub.EndDate = payload.EndDate
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.StartDate != nil {
		pub.StartDate = payload.StartDate
	}
	if payload.URL != nil {
		pub.URL = payload.URL
	}
	return &pub
}

// CreateEventPayload is the event create action payload.
type CreateEventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      string     `json:"name" xml:"name"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateEventPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if len(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true))
	}
	if payload.URL != nil {
		if len(*payload.URL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.url`, *payload.URL, len(*payload.URL), 5, true))
		}
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateEventContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteEventContext provides the event delete action context.
type DeleteEventContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	EventID  int
	TenantID int
}

// NewDeleteEventContext parses the incoming request URL and body, performs validations and creates the
// context used by the event controller delete action.
func NewDeleteEventContext(ctx context.Context, service *goa.Service) (*DeleteEventContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteEventContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteEventContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteEventContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListEventContext provides the event list action context.
type ListEventContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
}

// NewListEventContext parses the incoming request URL and body, performs validations and creates the
// context used by the event controller list action.
func NewListEventContext(ctx context.Context, service *goa.Service) (*ListEventContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListEventContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListEventContext) OK(r EventCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.event+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListEventContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowEventContext provides the event show action context.
type ShowEventContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	EventID  int
	TenantID int
}

// NewShowEventContext parses the incoming request URL and body, performs validations and creates the
// context used by the event controller show action.
func NewShowEventContext(ctx context.Context, service *goa.Service) (*ShowEventContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowEventContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowEventContext) OK(r *Event) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.event")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowEventContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateEventContext provides the event update action context.
type UpdateEventContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	EventID  int
	TenantID int
	Payload  *UpdateEventPayload
}

// NewUpdateEventContext parses the incoming request URL and body, performs validations and creates the
// context used by the event controller update action.
func NewUpdateEventContext(ctx context.Context, service *goa.Service) (*UpdateEventContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateEventContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// updateEventPayload is the event update action payload.
type updateEventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateEventPayload) Validate() (err error) {
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.URL != nil {
		if len(*payload.URL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.url`, *payload.URL, len(*payload.URL), 5, true))
		}
	}
	return err
}

// Publicize creates UpdateEventPayload from updateEventPayload
func (payload *updateEventPayload) Publicize() *UpdateEventPayload {
	var pub UpdateEventPayload
	if payload.EndDate != nil {
		pub.EndDate = payload.EndDate
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	if payload.StartDate != nil {
		pub.StartDate = payload.StartDate
	}
	if payload.URL != nil {
		pub.URL = payload.URL
	}
	return &pub
}

// UpdateEventPayload is the event update action payload.
type UpdateEventPayload struct {
	EndDate   *time.Time `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Name      *string    `json:"name,omitempty" xml:"name,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" xml:"start_date,omitempty"`
	URL       *string    `json:"url,omitempty" xml:"url,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateEventPayload) Validate() (err error) {
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	if payload.URL != nil {
		if len(*payload.URL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.url`, *payload.URL, len(*payload.URL), 5, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateEventContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateEventContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// StatusHealthzContext provides the healthz status action context.
type StatusHealthzContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewStatusHealthzContext parses the incoming request URL and body, performs validations and creates the
// context used by the healthz controller status action.
func NewStatusHealthzContext(ctx context.Context, service *goa.Service) (*StatusHealthzContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := StatusHealthzContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// CreatePresentationContext provides the presentation create action context.
type CreatePresentationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service   *goa.Service
	EventID   int
	SpeakerID int
	TenantID  int
	Payload   *CreatePresentationPayload
}

// NewCreatePresentationContext parses the incoming request URL and body, performs validations and creates the
// context used by the presentation controller create action.
func NewCreatePresentationContext(ctx context.Context, service *goa.Service) (*CreatePresentationContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreatePresentationContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// createPresentationPayload is the presentation create action payload.
type createPresentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createPresentationPayload) Validate() (err error) {
	if payload.Abstract == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "abstract"))
	}

	if payload.Abstract != nil {
		if len(*payload.Abstract) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 10, true))
		}
	}
	return err
}

// Publicize creates CreatePresentationPayload from createPresentationPayload
func (payload *createPresentationPayload) Publicize() *CreatePresentationPayload {
	var pub CreatePresentationPayload
	if payload.Abstract != nil {
		pub.Abstract = *payload.Abstract
	}
	if payload.Detail != nil {
		pub.Detail = payload.Detail
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	return &pub
}

// CreatePresentationPayload is the presentation create action payload.
type CreatePresentationPayload struct {
	Abstract string  `json:"abstract" xml:"abstract"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreatePresentationPayload) Validate() (err error) {
	if payload.Abstract == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "abstract"))
	}

	if len(payload.Abstract) < 10 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, payload.Abstract, len(payload.Abstract), 10, true))
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePresentationContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeletePresentationContext provides the presentation delete action context.
type DeletePresentationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service        *goa.Service
	EventID        int
	PresentationID int
	SpeakerID      int
	TenantID       int
}

// NewDeletePresentationContext parses the incoming request URL and body, performs validations and creates the
// context used by the presentation controller delete action.
func NewDeletePresentationContext(ctx context.Context, service *goa.Service) (*DeletePresentationContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeletePresentationContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramPresentationID := req.Params["presentationID"]
	if len(paramPresentationID) > 0 {
		rawPresentationID := paramPresentationID[0]
		if presentationID, err2 := strconv.Atoi(rawPresentationID); err2 == nil {
			rctx.PresentationID = presentationID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("presentationID", rawPresentationID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeletePresentationContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeletePresentationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListPresentationContext provides the presentation list action context.
type ListPresentationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service   *goa.Service
	EventID   int
	SpeakerID int
	TenantID  int
}

// NewListPresentationContext parses the incoming request URL and body, performs validations and creates the
// context used by the presentation controller list action.
func NewListPresentationContext(ctx context.Context, service *goa.Service) (*ListPresentationContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListPresentationContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OKAdmin sends a HTTP response with status code 200.
func (ctx *ListPresentationContext) OKAdmin(r PresentationAdminCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.presentation+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPresentationContext) OK(r PresentationCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.presentation+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListPresentationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowPresentationContext provides the presentation show action context.
type ShowPresentationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service        *goa.Service
	EventID        int
	PresentationID int
	SpeakerID      int
	TenantID       int
}

// NewShowPresentationContext parses the incoming request URL and body, performs validations and creates the
// context used by the presentation controller show action.
func NewShowPresentationContext(ctx context.Context, service *goa.Service) (*ShowPresentationContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowPresentationContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramPresentationID := req.Params["presentationID"]
	if len(paramPresentationID) > 0 {
		rawPresentationID := paramPresentationID[0]
		if presentationID, err2 := strconv.Atoi(rawPresentationID); err2 == nil {
			rctx.PresentationID = presentationID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("presentationID", rawPresentationID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OKAdmin sends a HTTP response with status code 200.
func (ctx *ShowPresentationContext) OKAdmin(r *PresentationAdmin) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.presentation")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPresentationContext) OK(r *Presentation) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.presentation")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPresentationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdatePresentationContext provides the presentation update action context.
type UpdatePresentationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service        *goa.Service
	EventID        int
	PresentationID int
	SpeakerID      int
	TenantID       int
	Payload        *UpdatePresentationPayload
}

// NewUpdatePresentationContext parses the incoming request URL and body, performs validations and creates the
// context used by the presentation controller update action.
func NewUpdatePresentationContext(ctx context.Context, service *goa.Service) (*UpdatePresentationContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdatePresentationContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramPresentationID := req.Params["presentationID"]
	if len(paramPresentationID) > 0 {
		rawPresentationID := paramPresentationID[0]
		if presentationID, err2 := strconv.Atoi(rawPresentationID); err2 == nil {
			rctx.PresentationID = presentationID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("presentationID", rawPresentationID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// updatePresentationPayload is the presentation update action payload.
type updatePresentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updatePresentationPayload) Validate() (err error) {
	if payload.Abstract != nil {
		if len(*payload.Abstract) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 10, true))
		}
	}
	return err
}

// Publicize creates UpdatePresentationPayload from updatePresentationPayload
func (payload *updatePresentationPayload) Publicize() *UpdatePresentationPayload {
	var pub UpdatePresentationPayload
	if payload.Abstract != nil {
		pub.Abstract = payload.Abstract
	}
	if payload.Detail != nil {
		pub.Detail = payload.Detail
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	return &pub
}

// UpdatePresentationPayload is the presentation update action payload.
type UpdatePresentationPayload struct {
	Abstract *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail   *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdatePresentationPayload) Validate() (err error) {
	if payload.Abstract != nil {
		if len(*payload.Abstract) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.abstract`, *payload.Abstract, len(*payload.Abstract), 10, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdatePresentationContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdatePresentationContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateSeriesContext provides the series create action context.
type CreateSeriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	Payload  *CreateSeriesPayload
}

// NewCreateSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller create action.
func NewCreateSeriesContext(ctx context.Context, service *goa.Service) (*CreateSeriesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateSeriesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// createSeriesPayload is the series create action payload.
type createSeriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createSeriesPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	return err
}

// Publicize creates CreateSeriesPayload from createSeriesPayload
func (payload *createSeriesPayload) Publicize() *CreateSeriesPayload {
	var pub CreateSeriesPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateSeriesPayload is the series create action payload.
type CreateSeriesPayload struct {
	Name string `json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateSeriesPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if len(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true))
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSeriesContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteSeriesContext provides the series delete action context.
type DeleteSeriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	SeriesID int
	TenantID int
}

// NewDeleteSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller delete action.
func NewDeleteSeriesContext(ctx context.Context, service *goa.Service) (*DeleteSeriesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteSeriesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramSeriesID := req.Params["seriesID"]
	if len(paramSeriesID) > 0 {
		rawSeriesID := paramSeriesID[0]
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			rctx.SeriesID = seriesID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteSeriesContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteSeriesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListSeriesContext provides the series list action context.
type ListSeriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
}

// NewListSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller list action.
func NewListSeriesContext(ctx context.Context, service *goa.Service) (*ListSeriesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListSeriesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSeriesContext) OK(r SeriesCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.series+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListSeriesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowSeriesContext provides the series show action context.
type ShowSeriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	SeriesID int
	TenantID int
}

// NewShowSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller show action.
func NewShowSeriesContext(ctx context.Context, service *goa.Service) (*ShowSeriesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowSeriesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramSeriesID := req.Params["seriesID"]
	if len(paramSeriesID) > 0 {
		rawSeriesID := paramSeriesID[0]
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			rctx.SeriesID = seriesID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSeriesContext) OK(r *Series) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.series")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSeriesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateSeriesContext provides the series update action context.
type UpdateSeriesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	SeriesID int
	TenantID int
	Payload  *UpdateSeriesPayload
}

// NewUpdateSeriesContext parses the incoming request URL and body, performs validations and creates the
// context used by the series controller update action.
func NewUpdateSeriesContext(ctx context.Context, service *goa.Service) (*UpdateSeriesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateSeriesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramSeriesID := req.Params["seriesID"]
	if len(paramSeriesID) > 0 {
		rawSeriesID := paramSeriesID[0]
		if seriesID, err2 := strconv.Atoi(rawSeriesID); err2 == nil {
			rctx.SeriesID = seriesID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("seriesID", rawSeriesID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// updateSeriesPayload is the series update action payload.
type updateSeriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateSeriesPayload) Validate() (err error) {
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	return err
}

// Publicize creates UpdateSeriesPayload from updateSeriesPayload
func (payload *updateSeriesPayload) Publicize() *UpdateSeriesPayload {
	var pub UpdateSeriesPayload
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	return &pub
}

// UpdateSeriesPayload is the series update action payload.
type UpdateSeriesPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateSeriesPayload) Validate() (err error) {
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateSeriesContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSeriesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateSpeakerContext provides the speaker create action context.
type CreateSpeakerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	EventID  int
	TenantID int
	Payload  *CreateSpeakerPayload
}

// NewCreateSpeakerContext parses the incoming request URL and body, performs validations and creates the
// context used by the speaker controller create action.
func NewCreateSpeakerContext(ctx context.Context, service *goa.Service) (*CreateSpeakerContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateSpeakerContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// createSpeakerPayload is the speaker create action payload.
type createSpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createSpeakerPayload) Validate() (err error) {
	if payload.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}

	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.ImageURL != nil {
		if len(*payload.ImageURL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.image_url`, *payload.ImageURL, len(*payload.ImageURL), 5, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	return err
}

// Publicize creates CreateSpeakerPayload from createSpeakerPayload
func (payload *createSpeakerPayload) Publicize() *CreateSpeakerPayload {
	var pub CreateSpeakerPayload
	if payload.Bio != nil {
		pub.Bio = payload.Bio
	}
	if payload.FirstName != nil {
		pub.FirstName = *payload.FirstName
	}
	if payload.Github != nil {
		pub.Github = payload.Github
	}
	if payload.ImageURL != nil {
		pub.ImageURL = payload.ImageURL
	}
	if payload.LastName != nil {
		pub.LastName = *payload.LastName
	}
	if payload.Linkedin != nil {
		pub.Linkedin = payload.Linkedin
	}
	if payload.Twitter != nil {
		pub.Twitter = payload.Twitter
	}
	return &pub
}

// CreateSpeakerPayload is the speaker create action payload.
type CreateSpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName string  `json:"first_name" xml:"first_name"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  string  `json:"last_name" xml:"last_name"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateSpeakerPayload) Validate() (err error) {
	if payload.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}

	if len(payload.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, payload.FirstName, len(payload.FirstName), 2, true))
	}
	if payload.ImageURL != nil {
		if len(*payload.ImageURL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.image_url`, *payload.ImageURL, len(*payload.ImageURL), 5, true))
		}
	}
	if len(payload.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, payload.LastName, len(payload.LastName), 2, true))
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSpeakerContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteSpeakerContext provides the speaker delete action context.
type DeleteSpeakerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service   *goa.Service
	EventID   int
	SpeakerID int
	TenantID  int
}

// NewDeleteSpeakerContext parses the incoming request URL and body, performs validations and creates the
// context used by the speaker controller delete action.
func NewDeleteSpeakerContext(ctx context.Context, service *goa.Service) (*DeleteSpeakerContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteSpeakerContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteSpeakerContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteSpeakerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListSpeakerContext provides the speaker list action context.
type ListSpeakerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	EventID  int
	TenantID int
}

// NewListSpeakerContext parses the incoming request URL and body, performs validations and creates the
// context used by the speaker controller list action.
func NewListSpeakerContext(ctx context.Context, service *goa.Service) (*ListSpeakerContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListSpeakerContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OKAdmin sends a HTTP response with status code 200.
func (ctx *ListSpeakerContext) OKAdmin(r SpeakerAdminCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.speaker+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSpeakerContext) OK(r SpeakerCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.speaker+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListSpeakerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowSpeakerContext provides the speaker show action context.
type ShowSpeakerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service   *goa.Service
	EventID   int
	SpeakerID int
	TenantID  int
}

// NewShowSpeakerContext parses the incoming request URL and body, performs validations and creates the
// context used by the speaker controller show action.
func NewShowSpeakerContext(ctx context.Context, service *goa.Service) (*ShowSpeakerContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowSpeakerContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OKAdmin sends a HTTP response with status code 200.
func (ctx *ShowSpeakerContext) OKAdmin(r *SpeakerAdmin) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.speaker")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSpeakerContext) OK(r *Speaker) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.speaker")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowSpeakerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateSpeakerContext provides the speaker update action context.
type UpdateSpeakerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service   *goa.Service
	EventID   int
	SpeakerID int
	TenantID  int
	Payload   *UpdateSpeakerPayload
}

// NewUpdateSpeakerContext parses the incoming request URL and body, performs validations and creates the
// context used by the speaker controller update action.
func NewUpdateSpeakerContext(ctx context.Context, service *goa.Service) (*UpdateSpeakerContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateSpeakerContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramSpeakerID := req.Params["speakerID"]
	if len(paramSpeakerID) > 0 {
		rawSpeakerID := paramSpeakerID[0]
		if speakerID, err2 := strconv.Atoi(rawSpeakerID); err2 == nil {
			rctx.SpeakerID = speakerID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("speakerID", rawSpeakerID, "integer"))
		}
	}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// updateSpeakerPayload is the speaker update action payload.
type updateSpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateSpeakerPayload) Validate() (err error) {
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.ImageURL != nil {
		if len(*payload.ImageURL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.image_url`, *payload.ImageURL, len(*payload.ImageURL), 5, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	return err
}

// Publicize creates UpdateSpeakerPayload from updateSpeakerPayload
func (payload *updateSpeakerPayload) Publicize() *UpdateSpeakerPayload {
	var pub UpdateSpeakerPayload
	if payload.Bio != nil {
		pub.Bio = payload.Bio
	}
	if payload.FirstName != nil {
		pub.FirstName = payload.FirstName
	}
	if payload.Github != nil {
		pub.Github = payload.Github
	}
	if payload.ImageURL != nil {
		pub.ImageURL = payload.ImageURL
	}
	if payload.LastName != nil {
		pub.LastName = payload.LastName
	}
	if payload.Linkedin != nil {
		pub.Linkedin = payload.Linkedin
	}
	if payload.Twitter != nil {
		pub.Twitter = payload.Twitter
	}
	return &pub
}

// UpdateSpeakerPayload is the speaker update action payload.
type UpdateSpeakerPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	Github    *string `json:"github,omitempty" xml:"github,omitempty"`
	ImageURL  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	LastName  *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty" xml:"linkedin,omitempty"`
	Twitter   *string `json:"twitter,omitempty" xml:"twitter,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateSpeakerPayload) Validate() (err error) {
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.ImageURL != nil {
		if len(*payload.ImageURL) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.image_url`, *payload.ImageURL, len(*payload.ImageURL), 5, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateSpeakerContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSpeakerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateTenantContext provides the tenant create action context.
type CreateTenantContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *CreateTenantPayload
}

// NewCreateTenantContext parses the incoming request URL and body, performs validations and creates the
// context used by the tenant controller create action.
func NewCreateTenantContext(ctx context.Context, service *goa.Service) (*CreateTenantContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateTenantContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// createTenantPayload is the tenant create action payload.
type createTenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createTenantPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	return err
}

// Publicize creates CreateTenantPayload from createTenantPayload
func (payload *createTenantPayload) Publicize() *CreateTenantPayload {
	var pub CreateTenantPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateTenantPayload is the tenant create action payload.
type CreateTenantPayload struct {
	Name string `json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateTenantPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if len(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true))
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateTenantContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteTenantContext provides the tenant delete action context.
type DeleteTenantContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
}

// NewDeleteTenantContext parses the incoming request URL and body, performs validations and creates the
// context used by the tenant controller delete action.
func NewDeleteTenantContext(ctx context.Context, service *goa.Service) (*DeleteTenantContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteTenantContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteTenantContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteTenantContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListTenantContext provides the tenant list action context.
type ListTenantContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewListTenantContext parses the incoming request URL and body, performs validations and creates the
// context used by the tenant controller list action.
func NewListTenantContext(ctx context.Context, service *goa.Service) (*ListTenantContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListTenantContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListTenantContext) OK(r TenantCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.tenant+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListTenantContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowTenantContext provides the tenant show action context.
type ShowTenantContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
}

// NewShowTenantContext parses the incoming request URL and body, performs validations and creates the
// context used by the tenant controller show action.
func NewShowTenantContext(ctx context.Context, service *goa.Service) (*ShowTenantContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowTenantContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTenantContext) OK(r *Tenant) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.tenant")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowTenantContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateTenantContext provides the tenant update action context.
type UpdateTenantContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	Payload  *UpdateTenantPayload
}

// NewUpdateTenantContext parses the incoming request URL and body, performs validations and creates the
// context used by the tenant controller update action.
func NewUpdateTenantContext(ctx context.Context, service *goa.Service) (*UpdateTenantContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateTenantContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// updateTenantPayload is the tenant update action payload.
type updateTenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateTenantPayload) Validate() (err error) {
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	return err
}

// Publicize creates UpdateTenantPayload from updateTenantPayload
func (payload *updateTenantPayload) Publicize() *UpdateTenantPayload {
	var pub UpdateTenantPayload
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	return &pub
}

// UpdateTenantPayload is the tenant update action payload.
type UpdateTenantPayload struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateTenantPayload) Validate() (err error) {
	if payload.Name != nil {
		if len(*payload.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, len(*payload.Name), 2, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateTenantContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateTenantContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// BootstrapUIContext provides the ui bootstrap action context.
type BootstrapUIContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewBootstrapUIContext parses the incoming request URL and body, performs validations and creates the
// context used by the ui controller bootstrap action.
func NewBootstrapUIContext(ctx context.Context, service *goa.Service) (*BootstrapUIContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := BootstrapUIContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *BootstrapUIContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	Payload  *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// createUserPayload is the user create action payload.
type createUserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.Role == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "role"))
	}

	if payload.Email != nil {
		if len(*payload.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, *payload.Email, len(*payload.Email), 2, true))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	if payload.Password != nil {
		if len(*payload.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, len(*payload.Password), 8, true))
		}
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.FirstName != nil {
		pub.FirstName = *payload.FirstName
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.LastName != nil {
		pub.LastName = *payload.LastName
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.Role != nil {
		pub.Role = *payload.Role
	}
	if payload.Validated != nil {
		pub.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		pub.ValidationCode = payload.ValidationCode
	}
	return &pub
}

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	Email          string  `json:"email" xml:"email"`
	FirstName      string  `json:"first_name" xml:"first_name"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       string  `json:"last_name" xml:"last_name"`
	Password       string  `json:"password" xml:"password"`
	Role           string  `json:"role" xml:"role"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "first_name"))
	}
	if payload.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "last_name"))
	}
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "role"))
	}

	if len(payload.Email) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, payload.Email, len(payload.Email), 2, true))
	}
	if len(payload.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, payload.FirstName, len(payload.FirstName), 2, true))
	}
	if len(payload.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, payload.LastName, len(payload.LastName), 2, true))
	}
	if len(payload.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, payload.Password, len(payload.Password), 8, true))
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	UserID   int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(ctx context.Context, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(ctx context.Context, service *goa.Service) (*ListUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ListUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OKTenant sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKTenant(r UserTenantCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	UserID   int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// OKTenant sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OKTenant(r *UserTenant) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	TenantID int
	UserID   int
	Payload  *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramTenantID := req.Params["tenantID"]
	if len(paramTenantID) > 0 {
		rawTenantID := paramTenantID[0]
		if tenantID, err2 := strconv.Atoi(rawTenantID); err2 == nil {
			rctx.TenantID = tenantID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("tenantID", rawTenantID, "integer"))
		}
	}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// updateUserPayload is the user update action payload.
type updateUserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUserPayload) Validate() (err error) {
	if payload.Email != nil {
		if len(*payload.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, *payload.Email, len(*payload.Email), 2, true))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	if payload.Password != nil {
		if len(*payload.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, len(*payload.Password), 8, true))
		}
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (payload *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if payload.Email != nil {
		pub.Email = payload.Email
	}
	if payload.FirstName != nil {
		pub.FirstName = payload.FirstName
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.LastName != nil {
		pub.LastName = payload.LastName
	}
	if payload.Password != nil {
		pub.Password = payload.Password
	}
	if payload.Role != nil {
		pub.Role = payload.Role
	}
	if payload.Validated != nil {
		pub.Validated = payload.Validated
	}
	if payload.ValidationCode != nil {
		pub.ValidationCode = payload.ValidationCode
	}
	return &pub
}

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	Email          *string `json:"email,omitempty" xml:"email,omitempty"`
	FirstName      *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	ID             *int    `json:"id,omitempty" xml:"id,omitempty"`
	LastName       *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Validated      *bool   `json:"validated,omitempty" xml:"validated,omitempty"`
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUserPayload) Validate() (err error) {
	if payload.Email != nil {
		if len(*payload.Email) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.email`, *payload.Email, len(*payload.Email), 2, true))
		}
	}
	if payload.FirstName != nil {
		if len(*payload.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.first_name`, *payload.FirstName, len(*payload.FirstName), 2, true))
		}
	}
	if payload.LastName != nil {
		if len(*payload.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.last_name`, *payload.LastName, len(*payload.LastName), 2, true))
		}
	}
	if payload.Password != nil {
		if len(*payload.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, len(*payload.Password), 8, true))
		}
	}
	if payload.ValidationCode != nil {
		if len(*payload.ValidationCode) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.validation_code`, *payload.ValidationCode, len(*payload.ValidationCode), 8, true))
		}
	}
	return err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ValidateValidateContext provides the validate validate action context.
type ValidateValidateContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	UserID  string
	V       *string
}

// NewValidateValidateContext parses the incoming request URL and body, performs validations and creates the
// context used by the validate controller validate action.
func NewValidateValidateContext(ctx context.Context, service *goa.Service) (*ValidateValidateContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ValidateValidateContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	paramV := req.Params["v"]
	if len(paramV) > 0 {
		rawV := paramV[0]
		rctx.V = &rawV
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ValidateValidateContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ValidateValidateContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
