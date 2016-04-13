package main

import (
	"github.com/goadesign/goa"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/models"
	"github.com/jinzhu/gorm"
)

// UserController implements theuser resource.
type UserController struct {
	*goa.Controller
	e           *Env
	storagelist map[string]Storage
	storage     models.UserStorage
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, storagelist map[string]Storage, e *Env) *UserController {
	return &UserController{
		Controller:  service.NewController("UserController"),
		e:           e,
		storagelist: storagelist,
		storage:     storagelist["USERSTORAGE"].(models.UserStorage),
	}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	a := models.UserFromCreateUserPayload(ctx.Payload)
	ra, err := c.storage.Add(ctx.Context, a)
	if err != nil {
		return ctx.Service.Send(ctx, 500, err.Error())
	}
	ctx.ResponseData.Header().Set("Location", app.UserHref(ctx.TenantID, ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	err := c.storage.Delete(ctx.Context, ctx.UserID)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	res := c.storage.ListUser(ctx.Context, ctx.TenantID)
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	res, err := c.storage.OneUser(ctx.Context, ctx.UserID, ctx.TenantID)
	if err != nil && err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	err := c.storage.UpdateFromUpdateUserPayload(ctx.Context, ctx.Payload, ctx.UserID)
	if err != nil {
		goa.LogError(ctx, "updating user", err.Error())
		return ctx.Err()
	}
	return ctx.NoContent()
}
