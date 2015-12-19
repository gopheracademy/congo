package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gopheracademy/congo/app"
	"github.com/gopheracademy/congo/handlers"
	"github.com/gopheracademy/congo/models"
	"github.com/gopheracademy/congo/swagger"
	"github.com/gopheracademy/congo/util"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
	"github.com/raphael/goa"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/lock"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
	//	_ "gopkg.in/authboss.v0/remember"
)

func main() {

	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://arrakis:8080/auth/twitter/callback"),
		//		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), "http://localhost:3000/auth/gplus/callback"),
		//		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		//		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:3000/auth/linkedin/callback"),
	)
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
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// Mount "user" controller
	c3 := NewUserController(service, models.NewUserDB(db))
	app.MountUserController(service, c3)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	goarouter := service.HTTPHandler().(*httprouter.Router)

	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	// Make sure to put authboss's router somewhere
	r := mux.NewRouter()
	gob.Register(models.User{})
	r.Handle("/auth/{provider}/callback", handlers.Landing(db, renderer)).Methods("GET")

	r.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler)
	r.Handle("/login", handlers.Login(renderer)).Methods("GET")
	gothic.GetProviderName = provider
	// Make sure to put authboss's router somewhere
	r.Handle("/", handlers.Index(renderer)).Methods("GET")
	r.Handle("/profile", handlers.Profile(renderer)).Methods("GET")
	r.Handle("/users", handlers.Users(renderer)).Methods("GET")

	r.PathPrefix("/api").Handler(goarouter)

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", cfg.Port)
	n.Run(hostStr)
}
func provider(r *http.Request) (string, error) {

	vars := mux.Vars(r)
	provider := vars["provider"]
	if provider == "" {
		return "", errors.New("No Provider specified")
	}
	return provider, nil
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

var indexTemplate = `
<p><a href="/auth/twitter">Log in with Twitter</a></p>
<p><a href="/auth/gplus">Log in with GPlus</a></p>
<p><a href="/auth/github">Log in with Github</a></p>
`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
`
