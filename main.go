package main

import (
	"fmt"
	"log"
	"os"

	jg "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	mjwt "github.com/goadesign/goa/middleware/security/jwt"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/jwt"
	"github.com/gopheracademy/congo/swagger"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

// settings holds the congo configuration.
var settings Settings

// Settings is the struct that holds config information retrieved
// from the environment
type Settings struct {
	DatabaseHost     string `envconfig:"db_host"`
	DatabaseUsername string `envconfig:"db_username"`
	DatabasePassword string `envconfig:"db_password"`
	DatabaseName     string `envconfig:"db_name"`
	DatabasePort     int    `envconfig:"db_port"`
	MaxOpenConns     int    `envconfig:"max_open" default:"100"`
	MaxIdleConns     int    `envconfig:"max_idle" default:"5"`
	Debug            bool   `envconfig:"debug" default:"false"`
}

// Env is the struct that holds important things to be passed into controllers
type Env struct {
	DB gorm.DB
}

func main() {

	err := envconfig.Process("congo", &settings)
	if err != nil {
		log.Fatal(err)
	}

	env, err := envFromSettings()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Migrate(env)

	storageList := makeStorage(env)
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	spec := &jwt.Specification{
		AllowParam:       false,
		AuthOptions:      false,
		TTLMinutes:       60,
		Issuer:           "congo.gopheracademy.com",
		KeySigningMethod: jwt.RSA256,
		SigningKeyFunc:   privateKey,
		ValidationFunc:   pubKey,
	}

	tm := jwt.NewTokenManager(spec)

	pubkey, err := readPubKey()
	if err != nil {
		panic(err)
	}
	pk, err := jg.ParseRSAPublicKeyFromPEM(pubkey)
	if err != nil {
		panic(err)
	}
	app.ConfigureJWTSecurity(service, mjwt.New(pk, nil))
	app.ConfigurePasswordSecurity(service, BasicAuth(env))
	// Mount "adminuser" controller
	c := NewAdminuserController(service)
	app.MountAdminuserController(service, c)
	// Mount "auth" controller
	c2 := NewAuthController(service, env, tm, spec)
	app.MountAuthController(service, c2)
	// Mount "event" controller
	c3 := NewEventController(service, storageList, env)
	app.MountEventController(service, c3)
	// Mount "healthz" controller
	c4 := NewHealthzController(service)
	app.MountHealthzController(service, c4)
	// Mount "series" controller
	c5 := NewSeriesController(service)
	app.MountSeriesController(service, c5)
	// Mount "tenant" controller
	c6 := NewTenantController(service, storageList, env)
	app.MountTenantController(service, c6)
	// Mount "user" controller
	c7 := NewUserController(service, storageList, env)
	app.MountUserController(service, c7)
	// Mount "validate" controller
	c8 := NewValidateController(service)
	app.MountValidateController(service, c8)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	service.ListenAndServe(":8080")
}
