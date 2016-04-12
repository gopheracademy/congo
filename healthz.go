package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// HealthzController implements the healthz resource.
type HealthzController struct {
	*goa.Controller
}

// NewHealthzController creates a healthz controller.
func NewHealthzController(service *goa.Service) *HealthzController {
	return &HealthzController{Controller: service.NewController("HealthzController")}
}

// Status runs the status action.
func (c *HealthzController) Status(ctx *app.StatusHealthzContext) error {
	// TBD: implement
	return nil
}
