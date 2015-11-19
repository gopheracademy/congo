package main

import (
	"github.com/bketelsen/congo/app"
	"github.com/bketelsen/congo/schema"
	"github.com/bketelsen/congo/swagger"
	"github.com/raphael/goa"
)

func main() {
	// Create service
	api := goa.NewGraceful("API")

	// Setup middleware
	api.Use(goa.RequestID())
	api.Use(goa.LogRequest())
	api.Use(goa.Recover())

	// Mount "account" controller
	c := NewAccountController()
	app.MountAccountController(api, c)
	// Mount "series" controller
	c2 := NewSeriesController()
	app.MountSeriesController(api, c2)
	// Mount "user" controller
	c3 := NewUserController()
	app.MountUserController(api, c3)

	// Mount Swagger spec provider controller
	swagger.MountController(api)

	// Mount json schema controller
	schema.MountController(api)
	// Start service, listen on port 8080
	api.ListenAndServe(":8080")
}
