package main

import (
	"fmt"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/raphael/goa"
)

// ProposalController implements the proposal resource.
type ProposalController struct {
	goa.Controller
	storage models.ProposalStorage
}

// NewProposalController creates a account controller.
func NewProposalController(service goa.Service, storage models.ProposalStorage) app.ProposalController {
	return &ProposalController{storage: storage, Controller: service.NewController("ProposalController")}
}

// Create runs the create action.
func (c *ProposalController) Create(ctx *app.CreateProposalContext) error {
	m, err := c.storage.Add(ctx, models.ProposalFromCreatePayload(ctx))
	if err != nil {
		return ctx.Err()
	}
	ctx.Header().Set("Location", app.ProposalHref(ctx.UserID, m.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *ProposalController) Delete(ctx *app.DeleteProposalContext) error {
	err := c.storage.Delete(ctx, ctx.ProposalID)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *ProposalController) List(ctx *app.ListProposalContext) error {
	res := c.storage.List(ctx)
	list := app.ProposalCollection{}
	for _, y := range res {
		fmt.Println(y)
		nm := y.ToApp()
		nm.Href = app.ProposalHref(y.UserID, y.ID)
		list = append(list, nm)
	}
	return ctx.OK(list)
}

// Show runs the show action.
func (c *ProposalController) Show(ctx *app.ShowProposalContext) error {
	res, err := c.storage.One(ctx, ctx.ProposalID)
	if err != nil {
		ctx.Error(err.Error())
	}
	m := res.ToApp()
	m.ID = int(res.ID)
	m.Href = app.ProposalHref(res.UserID, res.ID)

	return ctx.OK(m, "default")
}

// Update runs the update action.
func (c *ProposalController) Update(ctx *app.UpdateProposalContext) error {
	m := models.ProposalFromUpdatePayload(ctx)
	m.ID = ctx.UserID
	err := c.storage.Update(ctx, m)
	if err != nil {
		fmt.Println(err)
		return ctx.Err()
	}
	return ctx.NoContent()
}
