package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

// EventController implements theuser resource.
type EventController struct {
	*goa.Controller
	e           *Env
	storagelist map[string]Storage
	storage     models.EventStorage
}

// NewEventController creates a user controller.
func NewEventController(service *goa.Service, storagelist map[string]Storage, e *Env) *EventController {
	return &EventController{
		Controller:  service.NewController("EventController"),
		e:           e,
		storagelist: storagelist,
		storage:     storagelist["EVENTSTORAGE"].(models.EventStorage),
	}
}

// Create runs the create action.
func (c *EventController) Create(ctx *app.CreateEventContext) error {
	a := models.EventFromCreateEventPayload(ctx.Payload)
	ra, err := c.storage.Add(ctx.Context, a)
	if err != nil {
		return ctx.Service.Send(ctx, 500, err.Error())
	}
	ctx.ResponseData.Header().Set("Location", app.EventHref(ctx.TenantID, ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *EventController) Delete(ctx *app.DeleteEventContext) error {
	err := c.storage.Delete(ctx.Context, ctx.EventID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *EventController) List(ctx *app.ListEventContext) error {
	res := c.storage.ListEvent(ctx.Context, ctx.TenantID)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *EventController) Show(ctx *app.ShowEventContext) error {
	res, err := c.storage.OneEvent(ctx.Context, ctx.EventID, ctx.TenantID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *EventController) Update(ctx *app.UpdateEventContext) error {
	err := c.storage.UpdateFromUpdateEventPayload(ctx.Context, ctx.Payload, ctx.EventID)
	if err != nil {
		goa.LogError(ctx, "updating user", err.Error())
		return ctx.Err()
	}
	return ctx.NoContent()
}
