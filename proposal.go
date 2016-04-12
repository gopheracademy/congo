package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// ProposalController implements the proposal resource.
type ProposalController struct {
	*goa.Controller
}

// NewProposalController creates a proposal controller.
func NewProposalController(service *goa.Service) *ProposalController {
	return &ProposalController{Controller: service.NewController("proposal")}
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
	res := app.ProposalCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ProposalController) Show(ctx *app.ShowProposalContext) error {
	res := &app.Proposal{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ProposalController) Update(ctx *app.UpdateProposalContext) error {
	return nil
}
