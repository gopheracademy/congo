//************************************************************************//
// congo: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// AccountController is the controller interface for the Account actions.
type AccountController interface {
	goa.Controller
	Create(*CreateAccountContext) error
	Delete(*DeleteAccountContext) error
	List(*ListAccountContext) error
	Show(*ShowAccountContext) error
	Update(*UpdateAccountContext) error
}

// MountAccountController "mounts" a Account resource controller on the given service.
func MountAccountController(service goa.Service, ctrl AccountController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/api/accounts", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "Account", "action", "Create", "route", "POST /api/accounts")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/accounts/:accountID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Account", "action", "Delete", "route", "DELETE /api/accounts/:accountID")
	h = func(c *goa.Context) error {
		ctx, err := NewListAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/accounts", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Account", "action", "List", "route", "GET /api/accounts")
	h = func(c *goa.Context) error {
		ctx, err := NewShowAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Account", "action", "Show", "route", "GET /api/accounts/:accountID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PUT", "/api/accounts/:accountID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "Account", "action", "Update", "route", "PUT /api/accounts/:accountID")
}

// InstanceController is the controller interface for the Instance actions.
type InstanceController interface {
	goa.Controller
	Create(*CreateInstanceContext) error
	Delete(*DeleteInstanceContext) error
	List(*ListInstanceContext) error
	Show(*ShowInstanceContext) error
	Update(*UpdateInstanceContext) error
}

// MountInstanceController "mounts" a Instance resource controller on the given service.
func MountInstanceController(service goa.Service, ctrl InstanceController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateInstanceContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/api/accounts/:accountID/series/:seriesID/instances", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "Instance", "action", "Create", "route", "POST /api/accounts/:accountID/series/:seriesID/instances")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteInstanceContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/accounts/:accountID/series/:seriesID/instances/:instanceID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Instance", "action", "Delete", "route", "DELETE /api/accounts/:accountID/series/:seriesID/instances/:instanceID")
	h = func(c *goa.Context) error {
		ctx, err := NewListInstanceContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID/series/:seriesID/instances", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Instance", "action", "List", "route", "GET /api/accounts/:accountID/series/:seriesID/instances")
	h = func(c *goa.Context) error {
		ctx, err := NewShowInstanceContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID/series/:seriesID/instances/:instanceID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Instance", "action", "Show", "route", "GET /api/accounts/:accountID/series/:seriesID/instances/:instanceID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateInstanceContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/api/accounts/:accountID/series/:seriesID/instances/:instanceID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "Instance", "action", "Update", "route", "PATCH /api/accounts/:accountID/series/:seriesID/instances/:instanceID")
}

// SeriesController is the controller interface for the Series actions.
type SeriesController interface {
	goa.Controller
	Create(*CreateSeriesContext) error
	Delete(*DeleteSeriesContext) error
	List(*ListSeriesContext) error
	Show(*ShowSeriesContext) error
	Update(*UpdateSeriesContext) error
}

// MountSeriesController "mounts" a Series resource controller on the given service.
func MountSeriesController(service goa.Service, ctrl SeriesController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/api/accounts/:accountID/series", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "Series", "action", "Create", "route", "POST /api/accounts/:accountID/series")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/accounts/:accountID/series/:seriesID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Series", "action", "Delete", "route", "DELETE /api/accounts/:accountID/series/:seriesID")
	h = func(c *goa.Context) error {
		ctx, err := NewListSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID/series", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Series", "action", "List", "route", "GET /api/accounts/:accountID/series")
	h = func(c *goa.Context) error {
		ctx, err := NewShowSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID/series/:seriesID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Series", "action", "Show", "route", "GET /api/accounts/:accountID/series/:seriesID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/api/accounts/:accountID/series/:seriesID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "Series", "action", "Update", "route", "PATCH /api/accounts/:accountID/series/:seriesID")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Controller
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Show(*ShowUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service goa.Service, ctrl UserController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/api/accounts/:accountID/users", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "User", "action", "Create", "route", "POST /api/accounts/:accountID/users")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/accounts/:accountID/users/:userID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /api/accounts/:accountID/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewListUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID/users", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "User", "action", "List", "route", "GET /api/accounts/:accountID/users")
	h = func(c *goa.Context) error {
		ctx, err := NewShowUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/accounts/:accountID/users/:userID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "User", "action", "Show", "route", "GET /api/accounts/:accountID/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/api/accounts/:accountID/users/:userID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "User", "action", "Update", "route", "PATCH /api/accounts/:accountID/users/:userID")
}
