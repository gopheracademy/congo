package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/swagger"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())

	// Mount "auth" controller
	c := NewAuthController(service)
	app.MountAuthController(service, c)
	// Mount "proposal" controller
	c2 := NewProposalController(service)
	app.MountProposalController(service, c2)
	// Mount "review" controller
	c3 := NewReviewController(service)
	app.MountReviewController(service, c3)
	// Mount "ui" controller
	c4 := NewUIController(service)
	app.MountUIController(service, c4)
	// Mount "user" controller
	c5 := NewUserController(service)
	app.MountUserController(service, c5)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
