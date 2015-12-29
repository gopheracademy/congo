package main

import (
	"fmt"
	"strings"

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
	m, err := c.storage.Add(ctx, models.UserFromCreatePayload(ctx))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return ctx.Respond(400, []byte("Duplicate Registration"))
		} else {
			return ctx.Err()
		}
	}
	ctx.Header().Set("Location", app.UserHref(m.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	err := c.storage.Delete(ctx, ctx.UserID)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	res := c.storage.List(ctx)
	list := app.UserCollection{}
	for _, y := range res {
		fmt.Println(y)
		nm := y.ToApp()
		nm.Href = app.UserHref(y.ID)
		list = append(list, nm)
	}
	return ctx.OK(list)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	res, err := c.storage.One(ctx, ctx.UserID)
	if err != nil {
		ctx.Error(err.Error())
	}
	m := res.ToApp()
	m.ID = int(res.ID)
	m.Href = app.UserHref(res.ID)

	return ctx.OK(m, "default")
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	m := models.UserFromUpdatePayload(ctx)
	m.ID = ctx.UserID
	err := c.storage.Update(ctx, m)
	if err != nil {
		fmt.Println(err)
		return ctx.Err()
	}
	return ctx.NoContent()
}
