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
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	return err
}

// Publicize creates CreateEventPayload from createEventPayload
func (payload *createEventPayload) Publicize() *CreateEventPayload {
	var pub CreateEventPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateEventPayload is the event create action payload.
type CreateEventPayload struct {
	Name string `json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateEventPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if len(payload.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, len(payload.Name), 2, true))
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
