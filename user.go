package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/app/v1"
	"github.com/gopheracademy/congo/genmodels"
)

// UserController implements the account resource.
type UserController struct {
	goa.Controller
	storage genmodels.UserStorage
}

// NewUserController creates a account controller.
func NewUserController(service goa.Service, storage genmodels.UserStorage) v1.UserController {
	return &UserController{storage: storage, Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *v1.CreateUserContext) error {
	m, err := c.storage.Add(ctx, genmodels.UserFromV1CreatePayload(ctx))
	if err != nil {
		return ctx.Err()
	}
	ctx.Header().Set("Location", v1.UserHref("v1", m.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *v1.DeleteUserContext) error {
	err := c.storage.Delete(ctx, ctx.UserID)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *UserController) List(ctx *v1.ListUserContext) error {
	res := c.storage.List(ctx)
	list := app.UserCollection{}
	for _, y := range res {
		fmt.Println(y)
		nm := y.ToDefault()
		hr := v1.UserHref("v1", y.ID)
		nm.Href = &hr
		list = append(list, nm)
	}
	return ctx.OK(list)
}

// Show runs the show action.
func (c *UserController) Show(ctx *v1.ShowUserContext) error {
	res, err := c.storage.One(ctx, ctx.UserID)
	if err != nil {
		ctx.Error(err.Error())
	}
	m := res.ToDefault()
	i := int(res.ID)
	m.ID = &i
	hr := v1.UserHref("v1", res.ID)
	m.Href = &hr
	return ctx.OK(m, "default")
}

// Update runs the update action.
func (c *UserController) Update(ctx *v1.UpdateUserContext) error {
	m := user.UserFromV1UpdatePayload(ctx)
	m.ID = ctx.UserID
	err := c.storage.Update(ctx, m)
	if err != nil {
		fmt.Println(err)
		return ctx.Err()
	}
	return ctx.NoContent()
}
