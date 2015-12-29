package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/js"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/swagger"
	"github.com/gopheracademy/congo/util"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/raphael/goa"
	"github.com/raphael/goa-middleware/jwt"

	jg "github.com/dgrijalva/jwt-go"
)

func main() {

	cfg := util.Config{}
	if err := envconfig.Process("congo", &cfg); err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(goa.Recover())
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	spec := &jwt.Specification{
		AllowParam:       false,
		AuthOptions:      false,
		TTLMinutes:       60,
		Issuer:           "api.gopheracademy.com",
		KeySigningMethod: jwt.RSA256,
		SigningKeyFunc:   privateKey,
		ValidationFunc:   pubKey,
	}

	tm := jwt.NewTokenManager(spec)
	claims := make(map[string]interface{})
	claims["sub"] = "1"

	t, err := tm.Create(claims)
	fmt.Println(t)
	fmt.Println(err)

	a := NewAuthController(service, &db, tm, spec)
	app.MountAuthController(service, a)
	// Mount "user" controller
	c3 := NewUserController(service, models.NewUserDB(db))
	c3.Use(jwt.Middleware(spec))
	app.MountUserController(service, c3)

	js.MountController(service)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	service.ListenAndServe(fmt.Sprintf(":%d", cfg.Port))
}

func connectDB() (gorm.DB, error) {
	var db gorm.DB
	constr := fmt.Sprintf("user=%s host=%s port=%d dbname=%s password=%s sslmode=disable",
		"postgres",
		"127.0.0.1",
		5433,
		"postgres",
		"postgres")

	var err error
	db, err = gorm.Open("postgres", constr)
	if err == nil {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Proposal{})
		db.AutoMigrate(&models.Review{})
	}
	db.LogMode(true)
	return db, err
}

func pubKey(*jg.Token) (interface{}, error) {
	return ioutil.ReadFile("keys/gc.rsa.pub")
}

func privateKey() (interface{}, error) {
	return ioutil.ReadFile("keys/gc.rsa")
}
