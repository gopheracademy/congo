package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/jwt"
)

// NewAuthController creates a auth controller.
type AuthController struct {
	*goa.Controller
	e    *Env
	tm   *jwt.TokenManager
	spec *jwt.Specification
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, e *Env, tm *jwt.TokenManager, spec *jwt.Specification) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController"),
		e:    e,
		tm:   tm,
		spec: spec,
	}
}

// Refresh runs the refresh action.
func (c *AuthController) Refresh(ctx *app.RefreshAuthContext) error {
	// TBD: implement
	return nil
}

// Token runs the token action.
func (c *AuthController) Token(ctx *app.TokenAuthContext) error {
	// TBD: implement
	return nil
}
