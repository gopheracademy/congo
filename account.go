package main

import (
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/raphael/goa"
)

// AccountController implements the account resource.
type AccountController struct {
	goa.Controller
	storage models.AccountStorage
}

// NewAccountController creates a account controller.
func NewAccountController(service goa.Service, storage models.AccountStorage) app.AccountController {
	return &AccountController{storage: storage, Controller: service.NewController("AccountController")}
}

// Create runs the create action.
func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	m, err := c.storage.Add(ctx)
	if err != nil {
		return ctx.Err()
	}
	ctx.Header().Set("Location", app.AccountHref(m.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	err := c.storage.Delete(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	res := c.storage.List(ctx)
	list := app.AccountCollection{}
	for _, m := range res {
		a := m.ToApp()
		a.ID = int(m.ID)
		a.Href = app.AccountHref(a.ID)
		list = append(list, a)
	}
	return ctx.OK(list)
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	res, _ := c.storage.Get(ctx)
	m := res.ToApp()
	m.ID = int(res.ID)
	m.Href = app.AccountHref(res.ID)

	return ctx.OK(m, "default")
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	err := c.storage.Update(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}
