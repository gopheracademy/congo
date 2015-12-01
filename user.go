package main

import (
	"github.com/gopheracademy/congo/app"
	"github.com/raphael/goa"
)

// UserController implements the user resource.
type UserController struct {
	goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service goa.Service) app.UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	return nil
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	res := app.UserCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	res := &app.User{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	return nil
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	return nil
}
