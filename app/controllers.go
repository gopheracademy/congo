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

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Controller
	Callback(*CallbackAuthContext) error
	Refresh(*RefreshAuthContext) error
	Token(*TokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service goa.Service, ctrl AuthController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCallbackAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Callback(ctx)
	}
	router.Handle("GET", "/api/auth/:provider/callback", ctrl.NewHTTPRouterHandle("Callback", h))
	service.Info("mount", "ctrl", "Auth", "action", "Callback", "route", "GET /api/auth/:provider/callback")
	h = func(c *goa.Context) error {
		ctx, err := NewRefreshAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Refresh(ctx)
	}
	router.Handle("POST", "/api/auth/refresh", ctrl.NewHTTPRouterHandle("Refresh", h))
	service.Info("mount", "ctrl", "Auth", "action", "Refresh", "route", "POST /api/auth/refresh")
	h = func(c *goa.Context) error {
		ctx, err := NewTokenAuthContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Token(ctx)
	}
	router.Handle("POST", "/api/auth/token", ctrl.NewHTTPRouterHandle("Token", h))
	service.Info("mount", "ctrl", "Auth", "action", "Token", "route", "POST /api/auth/token")
}

// ProposalController is the controller interface for the Proposal actions.
type ProposalController interface {
	goa.Controller
	Create(*CreateProposalContext) error
	Delete(*DeleteProposalContext) error
	List(*ListProposalContext) error
	Show(*ShowProposalContext) error
	Update(*UpdateProposalContext) error
}

// MountProposalController "mounts" a Proposal resource controller on the given service.
func MountProposalController(service goa.Service, ctrl ProposalController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/api/users/:userID/proposals", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "Proposal", "action", "Create", "route", "POST /api/users/:userID/proposals")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/users/:userID/proposals/:proposalID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Proposal", "action", "Delete", "route", "DELETE /api/users/:userID/proposals/:proposalID")
	h = func(c *goa.Context) error {
		ctx, err := NewListProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/users/:userID/proposals", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Proposal", "action", "List", "route", "GET /api/users/:userID/proposals")
	h = func(c *goa.Context) error {
		ctx, err := NewShowProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/users/:userID/proposals/:proposalID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Proposal", "action", "Show", "route", "GET /api/users/:userID/proposals/:proposalID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/api/users/:userID/proposals/:proposalID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "Proposal", "action", "Update", "route", "PATCH /api/users/:userID/proposals/:proposalID")
}

// ReviewController is the controller interface for the Review actions.
type ReviewController interface {
	goa.Controller
	Create(*CreateReviewContext) error
	Delete(*DeleteReviewContext) error
	List(*ListReviewContext) error
	Show(*ShowReviewContext) error
	Update(*UpdateReviewContext) error
}

// MountReviewController "mounts" a Review resource controller on the given service.
func MountReviewController(service goa.Service, ctrl ReviewController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/api/users/:userID/proposals/:proposalID/review", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "Review", "action", "Create", "route", "POST /api/users/:userID/proposals/:proposalID/review")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Review", "action", "Delete", "route", "DELETE /api/users/:userID/proposals/:proposalID/review/:reviewID")
	h = func(c *goa.Context) error {
		ctx, err := NewListReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/users/:userID/proposals/:proposalID/review", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Review", "action", "List", "route", "GET /api/users/:userID/proposals/:proposalID/review")
	h = func(c *goa.Context) error {
		ctx, err := NewShowReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Review", "action", "Show", "route", "GET /api/users/:userID/proposals/:proposalID/review/:reviewID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "Review", "action", "Update", "route", "PATCH /api/users/:userID/proposals/:proposalID/review/:reviewID")
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
	router.Handle("POST", "/api/users", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "User", "action", "Create", "route", "POST /api/users")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/api/users/:userID", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /api/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewListUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/api/users", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "User", "action", "List", "route", "GET /api/users")
	h = func(c *goa.Context) error {
		ctx, err := NewShowUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/api/users/:userID", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "User", "action", "Show", "route", "GET /api/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PATCH", "/api/users/:userID", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "User", "action", "Update", "route", "PATCH /api/users/:userID")
}
