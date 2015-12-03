package main

import (
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/raphael/goa"
)

// SeriesController implements the account resource.
type SeriesController struct {
	goa.Controller
	storage models.SeriesStorage
}

// NewSeriesController creates a account controller.
func NewSeriesController(service goa.Service, storage models.SeriesStorage) app.SeriesController {
	return &SeriesController{storage: storage, Controller: service.NewController("SeriesController")}
} // Create runs the create action.
func (c *SeriesController) Create(ctx *app.CreateSeriesContext) error {
	_, err := c.storage.Add(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.Created()
}

// List runs the list action.
func (c *SeriesController) List(ctx *app.ListSeriesContext) error {
	res := c.storage.List(ctx)
	list := app.SeriesCollection{}
	for _, m := range res {

		list = append(list, m.ToApp())
	}
	return ctx.OK(list, "default")
}

// Show runs the show action.
func (c *SeriesController) Show(ctx *app.ShowSeriesContext) error {
	res := &app.Series{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *SeriesController) Update(ctx *app.UpdateSeriesContext) error {
	err := c.storage.Update(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// Delete runs the delete action.
func (c *SeriesController) Delete(ctx *app.DeleteSeriesContext) error {
	err := c.storage.Delete(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}
