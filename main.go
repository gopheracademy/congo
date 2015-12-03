package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/handlers"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/swagger"
	"github.com/gopheracademy/congo/util"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/raphael/goa"
)

func main() {

	cfg := util.Config{}
	if err := envconfig.Process("congo", &cfg); err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	env, err := cfg.Env()
	if err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}
	dev := env == util.EnvDev
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(goa.Recover())

	// Mount "account" controller
	c := NewAccountController(service, models.NewMockAccountStorage())
	app.MountAccountController(service, c)
	// Mount "series" controller
	c2 := NewSeriesController(service, models.NewMockSeriesStorage())
	app.MountSeriesController(service, c2)
	// Mount "user" controller
	c3 := NewUserController(service, models.NewMockUserStorage())
	app.MountUserController(service, c3)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	goarouter := service.HTTPHandler().(*httprouter.Router)

	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	r := mux.NewRouter()
	r.Handle("/", handlers.Index(renderer)).Methods("GET")
	r.Handle("/accounts/{accountID}/users", handlers.Users(renderer)).Methods("GET")
	r.Handle("/accounts", handlers.Accounts(renderer)).Methods("GET")

	r.Handle("/api/{_dummy:.*}", goarouter)
	n := negroni.Classic()

	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", cfg.Port)
	n.Run(hostStr)
}
