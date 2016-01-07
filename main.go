package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/app/v1"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/swagger"
	"github.com/gopheracademy/congo/util"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
	"github.com/raphael/goa"
	//	_ "gopkg.in/authboss.v0/remember"
	jg "github.com/dgrijalva/jwt-go"
	"github.com/raphael/goa-middleware/cors"
	"github.com/raphael/goa-middleware/jwt"
)

// JWT specification
var jwtSpec = &jwt.Specification{
	AllowParam:       false,
	AuthOptions:      false,
	TTLMinutes:       60,
	Issuer:           "api.gopheracademy.com",
	KeySigningMethod: jwt.RSA256,
	SigningKeyFunc:   privateKey,
	ValidationFunc:   pubKey,
}

func main() {

	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://arrakis:8080/api/auth/twitter/callback"),
		//		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), "http://localhost:3000/auth/gplus/callback"),
		//		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		//		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:3000/auth/linkedin/callback"),
	)
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

	// cors specification
	cspec, err := cors.New(func() {
		cors.Origin("*", func() {
			cors.Resource("/api", func() {
				cors.Headers("Accept", "Content-Type", "Origin", "Authorization")
				cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
				cors.MaxAge(600)
				cors.Credentials(true)
				cors.Vary("Http-Origin")
			})
		})
	})
	if err != nil {
		panic(err)
	}
	// mount the cors controller
	service.Use(cors.Middleware(cspec))
	cors.MountPreflightController(service, cspec)

	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	tm := jwt.NewTokenManager(jwtSpec)
	claims := make(map[string]interface{})
	claims["sub"] = "1"

	t, err := tm.Create(claims)
	// a token to get in through the backdoor printed to standard out
	fmt.Println(t)

	a := NewAuthController(service, &db, tm, jwtSpec)
	app.MountAuthController(service, a)
	// Mount "user" controller
	c3 := NewUserController(service, models.NewUserDB(db))
	v1.MountUserController(service, c3)
	c3.Use(jwt.Middleware(jwtSpec))
	c4 := NewProposalController(service, models.NewProposalDB(db))
	v1.MountProposalController(service, c4)
	c4.Use(jwt.Middleware(jwtSpec))
	ui := NewUIController(service, &db)
	app.MountUiController(service, ui)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	gothic.GetProviderName = provider

	hostStr := fmt.Sprintf(":%d", cfg.Port)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	service.ServeFiles("/assets/*filepath", filepath.Join(dir, "public"))
	service.ListenAndServe(hostStr)
}

func provider(r *http.Request) (string, error) {
	return "twitter", nil
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
	}
	db.LogMode(true)
	return db, err
}

func pubKey(*jg.Token) (interface{}, error) {
	//TODO(BJK) make the keys an ENV var
	f, err := os.Open("keys/gc.rsa.pub")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	return ioutil.ReadFile("keys/gc.rsa.pub")
}

func privateKey() (interface{}, error) {
	//TODO(BJK) make the keys an ENV var
	f, err := os.Open("keys/gc.rsa")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	return ioutil.ReadFile("keys/gc.rsa")
}
