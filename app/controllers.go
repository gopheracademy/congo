//************************************************************************//
// congo: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=/home/bketelsen/src/github.com/bketelsen/congo
// --design=github.com/bketelsen/congo/design
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
	Create(*CreateAccountContext) error
	Delete(*DeleteAccountContext) error
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
	router.Handle("POST", "/congo/accounts", goa.NewHTTPRouterHandle(service, "Account", "Create", h))
	service.Info("mount", "ctrl", "Account", "action", "Create", "route", "POST /congo/accounts")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/congo/accounts/:accountID", goa.NewHTTPRouterHandle(service, "Account", "Delete", h))
	service.Info("mount", "ctrl", "Account", "action", "Delete", "route", "DELETE /congo/accounts/:accountID")
	h = func(c *goa.Context) error {
		ctx, err := NewShowAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/congo/accounts/:accountID", goa.NewHTTPRouterHandle(service, "Account", "Show", h))
	service.Info("mount", "ctrl", "Account", "action", "Show", "route", "GET /congo/accounts/:accountID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateAccountContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PUT", "/congo/accounts/:accountID", goa.NewHTTPRouterHandle(service, "Account", "Update", h))
	service.Info("mount", "ctrl", "Account", "action", "Update", "route", "PUT /congo/accounts/:accountID")
}

// SeriesController is the controller interface for the Series actions.
type SeriesController interface {
	Create(*CreateSeriesContext) error
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
	router.Handle("POST", "/congo/accounts/:accountID/series", goa.NewHTTPRouterHandle(service, "Series", "Create", h))
	service.Info("mount", "ctrl", "Series", "action", "Create", "route", "POST /congo/accounts/:accountID/series")
	h = func(c *goa.Context) error {
		ctx, err := NewListSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/congo/accounts/:accountID/series", goa.NewHTTPRouterHandle(service, "Series", "List", h))
	service.Info("mount", "ctrl", "Series", "action", "List", "route", "GET /congo/accounts/:accountID/series")
	h = func(c *goa.Context) error {
		ctx, err := NewShowSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/congo/accounts/:accountID/series/:seriesID", goa.NewHTTPRouterHandle(service, "Series", "Show", h))
	service.Info("mount", "ctrl", "Series", "action", "Show", "route", "GET /congo/accounts/:accountID/series/:seriesID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateSeriesContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/congo/accounts/:accountID/series/:seriesID", goa.NewHTTPRouterHandle(service, "Series", "Update", h))
	service.Info("mount", "ctrl", "Series", "action", "Update", "route", "PATCH /congo/accounts/:accountID/series/:seriesID")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	Create(*CreateUserContext) error
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
	router.Handle("POST", "/congo/accounts/:accountID/users", goa.NewHTTPRouterHandle(service, "User", "Create", h))
	service.Info("mount", "ctrl", "User", "action", "Create", "route", "POST /congo/accounts/:accountID/users")
	h = func(c *goa.Context) error {
		ctx, err := NewListUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/congo/accounts/:accountID/users", goa.NewHTTPRouterHandle(service, "User", "List", h))
	service.Info("mount", "ctrl", "User", "action", "List", "route", "GET /congo/accounts/:accountID/users")
	h = func(c *goa.Context) error {
		ctx, err := NewShowUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/congo/accounts/:accountID/users/:userID", goa.NewHTTPRouterHandle(service, "User", "Show", h))
	service.Info("mount", "ctrl", "User", "action", "Show", "route", "GET /congo/accounts/:accountID/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/congo/accounts/:accountID/users/:userID", goa.NewHTTPRouterHandle(service, "User", "Update", h))
	service.Info("mount", "ctrl", "User", "action", "Update", "route", "PATCH /congo/accounts/:accountID/users/:userID")
}
