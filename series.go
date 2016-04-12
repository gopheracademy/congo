package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// SeriesController implements the series resource.
type SeriesController struct {
	*goa.Controller
}

// NewSeriesController creates a series controller.
func NewSeriesController(service *goa.Service) *SeriesController {
	return &SeriesController{Controller: service.NewController("SeriesController")}
}

// Create runs the create action.
func (c *SeriesController) Create(ctx *app.CreateSeriesContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *SeriesController) Delete(ctx *app.DeleteSeriesContext) error {
	// TBD: implement
	return nil
}

// List runs the list action.
func (c *SeriesController) List(ctx *app.ListSeriesContext) error {
	// TBD: implement
	res := app.SeriesCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *SeriesController) Show(ctx *app.ShowSeriesContext) error {
	// TBD: implement
	res := &app.Series{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *SeriesController) Update(ctx *app.UpdateSeriesContext) error {
	// TBD: implement
	return nil
}
