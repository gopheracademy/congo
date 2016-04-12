package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// TenantController implements the tenant resource.
type TenantController struct {
	*goa.Controller
}

// NewTenantController creates a tenant controller.
func NewTenantController(service *goa.Service) *TenantController {
	return &TenantController{Controller: service.NewController("TenantController")}
}

// Create runs the create action.
func (c *TenantController) Create(ctx *app.CreateTenantContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *TenantController) Delete(ctx *app.DeleteTenantContext) error {
	// TBD: implement
	return nil
}

// List runs the list action.
func (c *TenantController) List(ctx *app.ListTenantContext) error {
	// TBD: implement
	res := app.TenantCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TenantController) Show(ctx *app.ShowTenantContext) error {
	// TBD: implement
	res := &app.Tenant{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *TenantController) Update(ctx *app.UpdateTenantContext) error {
	// TBD: implement
	return nil
}
