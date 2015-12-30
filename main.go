package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/js"
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
		cors.Origin("https://cfp.gophercon.com", func() {
			cors.Resource("/api", func() {
				cors.Headers("Authorization")
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
	// a token to get in through the backdoor printed to standard out
	fmt.Println(t)

	a := NewAuthController(service, &db, tm, spec)
	app.MountAuthController(service, a)
	// Mount "user" controller
	c3 := NewUserController(service, models.NewUserDB(db))
	app.MountUserController(service, c3)
	c3.Use(jwt.Middleware(spec))
	c4 := NewProposalController(service, models.NewProposalDB(db))
	app.MountProposalController(service, c4)
	c4.Use(jwt.Middleware(spec))

	js.MountController(service)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	gothic.GetProviderName = provider

	hostStr := fmt.Sprintf(":%d", cfg.Port)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	service.ServeFiles("/assets/public/*filepath", http.Dir(filepath.Join(dir, "assets", "public")))
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
	return ioutil.ReadFile("keys/gc.rsa.pub")
}

func privateKey() (interface{}, error) {
	return ioutil.ReadFile("keys/gc.rsa")
}
