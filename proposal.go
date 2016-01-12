package main

import (
	"fmt"

	"github.com/gopheracademy/congo/app/v1"
	"github.com/gopheracademy/congo/models/proposal"
	"github.com/raphael/goa"
)

// ProposalController implements the proposal resource.
type ProposalController struct {
	goa.Controller
	storage proposal.ProposalStorage
}

// NewProposalController creates a account controller.
func NewProposalController(service goa.Service, storage proposal.ProposalStorage) v1.ProposalController {
	return &ProposalController{storage: storage, Controller: service.NewController("ProposalController")}
}

// Create runs the create action.
func (c *ProposalController) Create(ctx *v1.CreateProposalContext) error {
	m, err := c.storage.Add(ctx, proposal.ProposalFromV1CreatePayload(ctx))
	if err != nil {
		return ctx.Err()
	}
	ctx.Header().Set("Location", v1.ProposalHref(ctx.UserID, m.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *ProposalController) Delete(ctx *v1.DeleteProposalContext) error {
	err := c.storage.Delete(ctx, ctx.ProposalID)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *ProposalController) List(ctx *v1.ListProposalContext) error {
	res := c.storage.List(ctx)
	list := v1.ProposalCollection{}
	for _, y := range res {
		fmt.Println(y)
		nm := y.ToV1()
		nm.Href = v1.ProposalHref(y.UserID, y.ID)
		list = append(list, nm)
	}
	return ctx.OK(list)
}

// Show runs the show action.
func (c *ProposalController) Show(ctx *v1.ShowProposalContext) error {
	res, err := c.storage.One(ctx, ctx.ProposalID)
	if err != nil {
		ctx.Error(err.Error())
	}
	m := res.ToV1()
	m.ID = int(res.ID)
	m.Href = v1.ProposalHref(res.UserID, res.ID)

	return ctx.OK(m, "default")
}

// Update runs the update action.
func (c *ProposalController) Update(ctx *v1.UpdateProposalContext) error {
	m := proposal.ProposalFromV1UpdatePayload(ctx)
	m.ID = ctx.UserID
	err := c.storage.Update(ctx, m)
	if err != nil {
		fmt.Println(err)
		return ctx.Err()
	}
	return ctx.NoContent()
}
