package main

import (
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/raphael/goa"
)

// UserController implements the account resource.
type UserController struct {
	goa.Controller
	storage models.UserStorage
}

// NewUserController creates a account controller.
func NewUserController(service goa.Service, storage models.UserStorage) app.UserController {
	return &UserController{storage: storage, Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	m, err := c.storage.Add(ctx)
	if err != nil {
		return ctx.Err()
	}
	ctx.Header().Set("Location", app.UserHref(ctx.AccountID, m.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	err := c.storage.Delete(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	res := c.storage.List(ctx)
	list := app.UserCollection{}
	for _, m := range res {
		a := m.ToApp()
		a.ID = int(m.ID)
		a.Href = app.UserHref(ctx.AccountID, a.ID)
		list = append(list, a)
	}
	return ctx.OK(list)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	res, _ := c.storage.Get(ctx)
	m := res.ToApp()
	m.ID = int(res.ID)
	m.Href = app.UserHref(ctx.AccountID, res.ID)

	return ctx.OK(m, "default")
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	err := c.storage.Update(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}
