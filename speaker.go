package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

// SpeakerController implements theuser resource.
type SpeakerController struct {
	*goa.Controller
	e           *Env
	storagelist map[string]Storage
	storage     models.SpeakerStorage
}

// NewSpeakerController creates a user controller.
func NewSpeakerController(service *goa.Service, storagelist map[string]Storage, e *Env) *SpeakerController {
	return &SpeakerController{
		Controller:  service.NewController("SpeakerController"),
		e:           e,
		storagelist: storagelist,
		storage:     storagelist["SPEAKERSTORAGE"].(models.SpeakerStorage),
	}
}

// Create runs the create action.
func (c *SpeakerController) Create(ctx *app.CreateSpeakerContext) error {
	a := models.SpeakerFromCreateSpeakerPayload(ctx.Payload)
	ra, err := c.storage.Add(ctx.Context, a)
	if err != nil {
		return ctx.Service.Send(ctx, 500, err.Error())
	}
	ctx.ResponseData.Header().Set("Location", app.SpeakerHref(ctx.TenantID, ctx.EventID, ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *SpeakerController) Delete(ctx *app.DeleteSpeakerContext) error {
	err := c.storage.Delete(ctx.Context, ctx.SpeakerID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *SpeakerController) List(ctx *app.ListSpeakerContext) error {
	res := c.storage.ListSpeaker(ctx.Context, ctx.TenantID)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *SpeakerController) Show(ctx *app.ShowSpeakerContext) error {
	res, err := c.storage.OneSpeaker(ctx.Context, ctx.SpeakerID, ctx.TenantID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *SpeakerController) Update(ctx *app.UpdateSpeakerContext) error {
	err := c.storage.UpdateFromUpdateSpeakerPayload(ctx.Context, ctx.Payload, ctx.SpeakerID)
	if err != nil {
		goa.LogError(ctx, "updating user", err.Error())
		return ctx.Err()
	}
	return ctx.NoContent()
}
