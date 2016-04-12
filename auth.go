package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
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
