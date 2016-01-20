//************************************************************************//
// API "congo": Application Controllers
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

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Controller
	Callback(*CallbackAuthContext) error
	Oauth(*OauthAuthContext) error
	Refresh(*RefreshAuthContext) error
	Token(*TokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service goa.Service, ctrl AuthController) {
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewCallbackAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Callback(ctx)
	}
	mux.Handle("GET", "/auth/:provider/callback", ctrl.HandleFunc("Callback", h, nil))
	service.Info("mount", "ctrl", "Auth", "action", "Callback", "route", "GET /auth/:provider/callback")
	h = func(c *goa.Context) error {
		ctx, err := NewOauthAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Oauth(ctx)
	}
	mux.Handle("GET", "/auth/:provider", ctrl.HandleFunc("Oauth", h, nil))
	service.Info("mount", "ctrl", "Auth", "action", "Oauth", "route", "GET /auth/:provider")
	h = func(c *goa.Context) error {
		ctx, err := NewRefreshAuthContext(c)
		ctx.Payload = ctx.RawPayload().(*RefreshAuthPayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Refresh(ctx)
	}
	mux.Handle("POST", "/auth/refresh", ctrl.HandleFunc("Refresh", h, unmarshalRefreshAuthPayload))
	service.Info("mount", "ctrl", "Auth", "action", "Refresh", "route", "POST /auth/refresh")
	h = func(c *goa.Context) error {
		ctx, err := NewTokenAuthContext(c)
		ctx.Payload = ctx.RawPayload().(*TokenAuthPayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Token(ctx)
	}
	mux.Handle("POST", "/auth/token", ctrl.HandleFunc("Token", h, unmarshalTokenAuthPayload))
	service.Info("mount", "ctrl", "Auth", "action", "Token", "route", "POST /auth/token")
}

// unmarshalRefreshAuthPayload unmarshals the request body.
func unmarshalRefreshAuthPayload(ctx *goa.Context) error {
	payload := &RefreshAuthPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// unmarshalTokenAuthPayload unmarshals the request body.
func unmarshalTokenAuthPayload(ctx *goa.Context) error {
	payload := &TokenAuthPayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}

// UiController is the controller interface for the Ui actions.
type UiController interface {
	goa.Controller
	Bootstrap(*BootstrapUiContext) error
}

// MountUiController "mounts" a Ui resource controller on the given service.
func MountUiController(service goa.Service, ctrl UiController) {
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewBootstrapUiContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Bootstrap(ctx)
	}
	mux.Handle("GET", "/", ctrl.HandleFunc("Bootstrap", h, nil))
	service.Info("mount", "ctrl", "Ui", "action", "Bootstrap", "route", "GET /")
}
