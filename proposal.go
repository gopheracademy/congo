package main

import (
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
	return nil
}

// Delete runs the delete action.
func (c *ProposalController) Delete(ctx *app.DeleteProposalContext) error {
	return nil
}

// List runs the list action.
func (c *ProposalController) List(ctx *app.ListProposalContext) error {
	return nil
}

// Show runs the show action.
func (c *ProposalController) Show(ctx *app.ShowProposalContext) error {
	res := &app.Proposal{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *ProposalController) Update(ctx *app.UpdateProposalContext) error {
	return nil
}
