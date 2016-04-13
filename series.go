package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

// SeriesController implements theuser resource.
type SeriesController struct {
	*goa.Controller
	e           *Env
	storagelist map[string]Storage
	storage     models.SeriesStorage
}

// NewSeriesController creates a user controller.
func NewSeriesController(service *goa.Service, storagelist map[string]Storage, e *Env) *SeriesController {
	return &SeriesController{
		Controller:  service.NewController("SeriesController"),
		e:           e,
		storagelist: storagelist,
		storage:     storagelist["SERIESSTORAGE"].(models.SeriesStorage),
	}
}

// Create runs the create action.
func (c *SeriesController) Create(ctx *app.CreateSeriesContext) error {
	a := models.SeriesFromCreateSeriesPayload(ctx.Payload)
	ra, err := c.storage.Add(ctx.Context, a)
	if err != nil {
		return ctx.Service.Send(ctx, 500, err.Error())
	}
	ctx.ResponseData.Header().Set("Location", app.SeriesHref(ctx.TenantID, ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *SeriesController) Delete(ctx *app.DeleteSeriesContext) error {
	err := c.storage.Delete(ctx.Context, ctx.SeriesID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *SeriesController) List(ctx *app.ListSeriesContext) error {
	res := c.storage.ListSeries(ctx.Context, ctx.TenantID)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *SeriesController) Show(ctx *app.ShowSeriesContext) error {
	res, err := c.storage.OneSeries(ctx.Context, ctx.SeriesID, ctx.TenantID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *SeriesController) Update(ctx *app.UpdateSeriesContext) error {
	err := c.storage.UpdateFromUpdateSeriesPayload(ctx.Context, ctx.Payload, ctx.SeriesID)
	if err != nil {
		goa.LogError(ctx, "updating user", err.Error())
		return ctx.Err()
	}
	return ctx.NoContent()
}
