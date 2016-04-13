package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

// TenantController implements theuser resource.
type TenantController struct {
	*goa.Controller
	e           *Env
	storagelist map[string]Storage
	storage     models.TenantStorage
}

// NewTenantController creates a tenant controller.
func NewTenantController(service *goa.Service, storagelist map[string]Storage, e *Env) *TenantController {
	return &TenantController{
		Controller:  service.NewController("TenantController"),
		e:           e,
		storagelist: storagelist,
		storage:     storagelist["TENANTSTORAGE"].(models.TenantStorage),
	}
}

// Create runs the create action.
func (c *TenantController) Create(ctx *app.CreateTenantContext) error {
	a := models.TenantFromCreateTenantPayload(ctx.Payload)
	ra, err := c.storage.Add(ctx.Context, a)
	if err != nil {
		return ctx.Service.Send(ctx, 500, err.Error())
	}
	ctx.ResponseData.Header().Set("Location", app.TenantHref(ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *TenantController) Delete(ctx *app.DeleteTenantContext) error {
	err := c.storage.Delete(ctx.Context, ctx.TenantID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *TenantController) List(ctx *app.ListTenantContext) error {
	res := c.storage.ListTenant(ctx.Context)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TenantController) Show(ctx *app.ShowTenantContext) error {
	res, err := c.storage.OneTenant(ctx.Context, ctx.TenantID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *TenantController) Update(ctx *app.UpdateTenantContext) error {
	err := c.storage.UpdateFromUpdateTenantPayload(ctx.Context, ctx.Payload, ctx.TenantID)
	if err != nil {
		goa.LogError(ctx, "updating tenant", err.Error())
		return ctx.Err()
	}
	return ctx.NoContent()
}
