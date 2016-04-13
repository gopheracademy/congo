package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// ValidateController implements the validate resource.
type ValidateController struct {
	*goa.Controller
}

// NewValidateController creates a validate controller.
func NewValidateController(service *goa.Service) *ValidateController {
	return &ValidateController{Controller: service.NewController("ValidateController")}
}

// Validate runs the validate action.
func (c *ValidateController) Validate(ctx *app.ValidateValidateContext) error {
	// TBD: implement
	return nil
}
