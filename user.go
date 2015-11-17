package main

import (
	"github.com/bketelsen/congo/app"
)

// UserController implements the user resource.
type UserController struct{}

// NewUserController creates a user controller.
func NewUserController() *UserController {
	return &UserController{}
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
