package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// AdminuserController implements the adminuser resource.
type AdminuserController struct {
	*goa.Controller
}

// NewAdminuserController creates a adminuser controller.
func NewAdminuserController(service *goa.Service) *AdminuserController {
	return &AdminuserController{Controller: service.NewController("AdminuserController")}
}

// Create runs the create action.
func (c *AdminuserController) Create(ctx *app.CreateAdminuserContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *AdminuserController) Delete(ctx *app.DeleteAdminuserContext) error {
	// TBD: implement
	return nil
}

// List runs the list action.
func (c *AdminuserController) List(ctx *app.ListAdminuserContext) error {
	// TBD: implement
	res := app.UserCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AdminuserController) Show(ctx *app.ShowAdminuserContext) error {
	// TBD: implement
	res := &app.User{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AdminuserController) Update(ctx *app.UpdateAdminuserContext) error {
	// TBD: implement
	return nil
}
