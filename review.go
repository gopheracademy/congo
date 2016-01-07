package main

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/raphael/goa"
)

// ReviewController implements the review resource.
type ReviewController struct {
	goa.Controller
}

// NewReviewController creates a review controller.
func NewReviewController(service goa.Service) v1.ReviewController {
	return &ReviewController{Controller: service.NewController("ReviewController")}
}

// Create runs the create action.
func (c *ReviewController) Create(ctx *v1.CreateReviewContext) error {
	return nil
}

// Delete runs the delete action.
func (c *ReviewController) Delete(ctx *v1.DeleteReviewContext) error {
	return nil
}

// List runs the list action.
func (c *ReviewController) List(ctx *v1.ListReviewContext) error {
	return nil
}

// Show runs the show action.
func (c *ReviewController) Show(ctx *v1.ShowReviewContext) error {
	res := &v1.Review{}
	return ctx.OK(res, "default")
}

// Update runs the update action.
func (c *ReviewController) Update(ctx *v1.UpdateReviewContext) error {
	return nil
}
