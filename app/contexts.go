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

import "github.com/goadesign/goa"

// CallbackAuthContext provides the auth callback action context.
type CallbackAuthContext struct {
	*goa.Context
	APIVersion string
	Provider   string
}

// NewCallbackAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller callback action.
func NewCallbackAuthContext(c *goa.Context) (*CallbackAuthContext, error) {
	var err error
	ctx := CallbackAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProvider := c.Get("provider")
	if rawProvider != "" {
		ctx.Provider = rawProvider
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CallbackAuthContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "text/html")
	return ctx.RespondBytes(200, resp)
}

// OauthAuthContext provides the auth oauth action context.
type OauthAuthContext struct {
	*goa.Context
	APIVersion string
	Provider   string
}

// NewOauthAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller oauth action.
func NewOauthAuthContext(c *goa.Context) (*OauthAuthContext, error) {
	var err error
	ctx := OauthAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	rawProvider := c.Get("provider")
	if rawProvider != "" {
		ctx.Provider = rawProvider
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *OauthAuthContext) OK(resp *Authorize) error {
	ctx.Header().Set("Content-Type", "application/vnd.authorize")
	return ctx.Respond(200, resp)
}

// RefreshAuthContext provides the auth refresh action context.
type RefreshAuthContext struct {
	*goa.Context
	APIVersion string
	Payload    *RefreshAuthPayload
}

// NewRefreshAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller refresh action.
func NewRefreshAuthContext(c *goa.Context) (*RefreshAuthContext, error) {
	var err error
	ctx := RefreshAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	return &ctx, err
}

// RefreshAuthPayload is the auth refresh action payload.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// Created sends a HTTP response with status code 201.
func (ctx *RefreshAuthContext) Created(resp *Authorize) error {
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.Respond(201, resp)
}

// TokenAuthContext provides the auth token action context.
type TokenAuthContext struct {
	*goa.Context
	APIVersion string
	Payload    *TokenAuthPayload
}

// NewTokenAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller token action.
func NewTokenAuthContext(c *goa.Context) (*TokenAuthContext, error) {
	var err error
	ctx := TokenAuthContext{Context: c}
	rawAPIVersion := c.Get("api_version")
	if rawAPIVersion != "" {
		ctx.APIVersion = rawAPIVersion
	}
	return &ctx, err
}

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application *string `json:"application,omitempty" xml:"application,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

// Created sends a HTTP response with status code 201.
func (ctx *TokenAuthContext) Created(resp *Authorize) error {
	ctx.Header().Set("Content-Type", "application/vnd.authorize+json")
	return ctx.Respond(201, resp)
}

// BootstrapUiContext provides the ui bootstrap action context.
type BootstrapUiContext struct {
	*goa.Context
}

// NewBootstrapUiContext parses the incoming request URL and body, performs validations and creates the
// context used by the ui controller bootstrap action.
func NewBootstrapUiContext(c *goa.Context) (*BootstrapUiContext, error) {
	var err error
	ctx := BootstrapUiContext{Context: c}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *BootstrapUiContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "text/html")
	return ctx.RespondBytes(200, resp)
}
