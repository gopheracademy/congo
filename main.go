package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/swagger"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "adminuser" controller
	c := NewAdminuserController(service)
	app.MountAdminuserController(service, c)
	// Mount "auth" controller
	c2 := NewAuthController(service)
	app.MountAuthController(service, c2)
	// Mount "event" controller
	c3 := NewEventController(service)
	app.MountEventController(service, c3)
	// Mount "healthz" controller
	c4 := NewHealthzController(service)
	app.MountHealthzController(service, c4)
	// Mount "series" controller
	c5 := NewSeriesController(service)
	app.MountSeriesController(service, c5)
	// Mount "tenant" controller
	c6 := NewTenantController(service)
	app.MountTenantController(service, c6)
	// Mount "user" controller
	c7 := NewUserController(service)
	app.MountUserController(service, c7)
	// Mount "validate" controller
	c8 := NewValidateController(service)
	app.MountValidateController(service, c8)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	service.ListenAndServe(":8080")
}
