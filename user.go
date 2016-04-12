package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// TBD: implement
	return nil
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// TBD: implement
	res := app.UserCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// TBD: implement
	res := &app.User{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// TBD: implement
	return nil
}
