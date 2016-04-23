package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

// PresentationController implements theuser resource.
type PresentationController struct {
	*goa.Controller
	e           *Env
	storagelist map[string]Storage
	storage     models.PresentationStorage
}

// NewPresentationController creates a user controller.
func NewPresentationController(service *goa.Service, storagelist map[string]Storage, e *Env) *PresentationController {
	return &PresentationController{
		Controller:  service.NewController("PresentationController"),
		e:           e,
		storagelist: storagelist,
		storage:     storagelist["PRESENTATIONSTORAGE"].(models.PresentationStorage),
	}
}

// Create runs the create action.
func (c *PresentationController) Create(ctx *app.CreatePresentationContext) error {
	a := models.PresentationFromCreatePresentationPayload(ctx.Payload)
	ra, err := c.storage.Add(ctx.Context, a)
	if err != nil {
		return ctx.Service.Send(ctx, 500, err.Error())
	}
	ctx.ResponseData.Header().Set("Location", app.PresentationHref(ctx.TenantID, ctx.EventID, ctx.SpeakerID, ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *PresentationController) Delete(ctx *app.DeletePresentationContext) error {
	err := c.storage.Delete(ctx.Context, ctx.PresentationID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *PresentationController) List(ctx *app.ListPresentationContext) error {
	res := c.storage.ListPresentation(ctx.Context, ctx.EventID, ctx.SpeakerID)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *PresentationController) Show(ctx *app.ShowPresentationContext) error {
	res, err := c.storage.OnePresentation(ctx.Context, ctx.PresentationID, ctx.EventID, ctx.SpeakerID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *PresentationController) Update(ctx *app.UpdatePresentationContext) error {
	err := c.storage.UpdateFromUpdatePresentationPayload(ctx.Context, ctx.Payload, ctx.PresentationID)
	if err != nil {
		goa.LogError(ctx, "updating user", err.Error())
		return ctx.Err()
	}
	return ctx.NoContent()
}
