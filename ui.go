package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// UIController implements the ui resource.
type UIController struct {
	*goa.Controller
}

// NewUIController creates a ui controller.
func NewUIController(service *goa.Service) *UIController {
	return &UIController{Controller: service.NewController("ui")}
}

// Bootstrap runs the bootstrap action.
func (c *UIController) Bootstrap(ctx *app.BootstrapUIContext) error {
	return nil
}
