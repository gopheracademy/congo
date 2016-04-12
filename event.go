package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
)

// EventController implements the event resource.
type EventController struct {
	*goa.Controller
}

// NewEventController creates a event controller.
func NewEventController(service *goa.Service) *EventController {
	return &EventController{Controller: service.NewController("EventController")}
}

// Create runs the create action.
func (c *EventController) Create(ctx *app.CreateEventContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *EventController) Delete(ctx *app.DeleteEventContext) error {
	// TBD: implement
	return nil
}

// List runs the list action.
func (c *EventController) List(ctx *app.ListEventContext) error {
	// TBD: implement
	res := app.EventCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *EventController) Show(ctx *app.ShowEventContext) error {
	// TBD: implement
	res := &app.Event{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *EventController) Update(ctx *app.UpdateEventContext) error {
	// TBD: implement
	return nil
}
