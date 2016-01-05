//************************************************************************//
// API "congo" version 1.0.0: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/gopheracademy/congo
// --design=github.com/gopheracademy/congo/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v100

import "github.com/raphael/goa"

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
	var h goa.Handler
	mux := service.ServeMux().Version("1.0.0")
	h = func(c *goa.Context) error {
		ctx, err := NewCreateProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/api/users/:userID/proposals", ctrl.HandleFunc("Create", h))
	service.Info("mount", "ctrl", "Proposal", "version", "1.0.0", "action", "Create", "route", "POST /api/users/:userID/proposals")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Delete", h))
	service.Info("mount", "ctrl", "Proposal", "version", "1.0.0", "action", "Delete", "route", "DELETE /api/users/:userID/proposals/:proposalID")
	h = func(c *goa.Context) error {
		ctx, err := NewListProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals", ctrl.HandleFunc("List", h))
	service.Info("mount", "ctrl", "Proposal", "version", "1.0.0", "action", "List", "route", "GET /api/users/:userID/proposals")
	h = func(c *goa.Context) error {
		ctx, err := NewShowProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Show", h))
	service.Info("mount", "ctrl", "Proposal", "version", "1.0.0", "action", "Show", "route", "GET /api/users/:userID/proposals/:proposalID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateProposalContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/api/users/:userID/proposals/:proposalID", ctrl.HandleFunc("Update", h))
	service.Info("mount", "ctrl", "Proposal", "version", "1.0.0", "action", "Update", "route", "PATCH /api/users/:userID/proposals/:proposalID")
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
	var h goa.Handler
	mux := service.ServeMux().Version("1.0.0")
	h = func(c *goa.Context) error {
		ctx, err := NewCreateReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/api/users/:userID/proposals/:proposalID/review", ctrl.HandleFunc("Create", h))
	service.Info("mount", "ctrl", "Review", "version", "1.0.0", "action", "Create", "route", "POST /api/users/:userID/proposals/:proposalID/review")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Delete", h))
	service.Info("mount", "ctrl", "Review", "version", "1.0.0", "action", "Delete", "route", "DELETE /api/users/:userID/proposals/:proposalID/review/:reviewID")
	h = func(c *goa.Context) error {
		ctx, err := NewListReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals/:proposalID/review", ctrl.HandleFunc("List", h))
	service.Info("mount", "ctrl", "Review", "version", "1.0.0", "action", "List", "route", "GET /api/users/:userID/proposals/:proposalID/review")
	h = func(c *goa.Context) error {
		ctx, err := NewShowReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Show", h))
	service.Info("mount", "ctrl", "Review", "version", "1.0.0", "action", "Show", "route", "GET /api/users/:userID/proposals/:proposalID/review/:reviewID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateReviewContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/api/users/:userID/proposals/:proposalID/review/:reviewID", ctrl.HandleFunc("Update", h))
	service.Info("mount", "ctrl", "Review", "version", "1.0.0", "action", "Update", "route", "PATCH /api/users/:userID/proposals/:proposalID/review/:reviewID")
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
	var h goa.Handler
	mux := service.ServeMux().Version("1.0.0")
	h = func(c *goa.Context) error {
		ctx, err := NewCreateUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	mux.Handle("POST", "/api/users", ctrl.HandleFunc("Create", h))
	service.Info("mount", "ctrl", "User", "version", "1.0.0", "action", "Create", "route", "POST /api/users")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	mux.Handle("DELETE", "/api/users/:userID", ctrl.HandleFunc("Delete", h))
	service.Info("mount", "ctrl", "User", "version", "1.0.0", "action", "Delete", "route", "DELETE /api/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewListUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	mux.Handle("GET", "/api/users", ctrl.HandleFunc("List", h))
	service.Info("mount", "ctrl", "User", "version", "1.0.0", "action", "List", "route", "GET /api/users")
	h = func(c *goa.Context) error {
		ctx, err := NewShowUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	mux.Handle("GET", "/api/users/:userID", ctrl.HandleFunc("Show", h))
	service.Info("mount", "ctrl", "User", "version", "1.0.0", "action", "Show", "route", "GET /api/users/:userID")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateUserContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	mux.Handle("PATCH", "/api/users/:userID", ctrl.HandleFunc("Update", h))
	service.Info("mount", "ctrl", "User", "version", "1.0.0", "action", "Update", "route", "PATCH /api/users/:userID")
}
