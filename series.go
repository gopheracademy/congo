package main

import (
	"github.com/gopheracademy/congo/app"
)

// SeriesController implements the series resource.
type SeriesController struct{}

// NewSeriesController creates a series controller.
func NewSeriesController() *SeriesController {
	return &SeriesController{}
}

// Create runs the create action.
func (c *SeriesController) Create(ctx *app.CreateSeriesContext) error {
	return nil
}

// List runs the list action.
func (c *SeriesController) List(ctx *app.ListSeriesContext) error {
	res := app.SeriesCollection{}
	return ctx.OK(res, "default")
}

// Show runs the show action.
func (c *SeriesController) Show(ctx *app.ShowSeriesContext) error {
	res := &app.Series{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *SeriesController) Update(ctx *app.UpdateSeriesContext) error {
	return nil
}
