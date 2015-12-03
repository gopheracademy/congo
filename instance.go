package main

import (
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/raphael/goa"
)

// InstanceController implements the account resource.
type InstanceController struct {
	goa.Controller
	storage models.InstanceStorage
}

// NewInstanceController creates a account controller.
func NewInstanceController(service goa.Service, storage models.InstanceStorage) app.InstanceController {
	return &InstanceController{storage: storage, Controller: service.NewController("InstanceController")}
} // Create runs the create action.
func (c *InstanceController) Create(ctx *app.CreateInstanceContext) error {
	_, err := c.storage.Add(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.Created()
}

// List runs the list action.
func (c *InstanceController) List(ctx *app.ListInstanceContext) error {
	res := c.storage.List(ctx)
	list := app.InstanceCollection{}
	for _, m := range res {

		list = append(list, m.ToApp())
	}
	return ctx.OK(list, "default")
}

// Show runs the show action.
func (c *InstanceController) Show(ctx *app.ShowInstanceContext) error {
	res := &app.Instance{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *InstanceController) Update(ctx *app.UpdateInstanceContext) error {
	err := c.storage.Update(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}

// Delete runs the delete action.
func (c *InstanceController) Delete(ctx *app.DeleteInstanceContext) error {
	err := c.storage.Delete(ctx)
	if err != nil {
		return ctx.Err()
	}
	return ctx.NoContent()
}
